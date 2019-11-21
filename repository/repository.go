package repository

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/ilamazuliaf/desaku/models"

	uuid "github.com/satori/go.uuid"
)

const (
	userLogin = `SELECT admin.uuid_person, admin.username, admin.password, admin.email, user_group.nama, admin.aktif
			FROM admin
			INNER JOIN user_group ON admin.user_id_group = user_group.id
			WHERE admin.username = ? AND admin.password = ?`
	personInsert = `insert person set uuid = ?, nokk = ? , nik = ?, 
			nama = ? , jk =?, tempat_lahir = ?, tanggal_lahir = ? , 
			anak_ke = ?, jum_saudara = ? , phone1 = ? , phone2 = ? ,
			id_pendidikan = ? , id_pekerjaan = ? , id_penghasilan = ? ,
			pembuat = ?`
	personSelectBase = `SELECT person.uuid, person.nokk, person.nik, person.nama, person.jk, person.tempat_lahir,
			person.tanggal_lahir, person.anak_ke, person.jum_saudara, person.phone1,
			person.phone2, pendidikan.id, pendidikan.nama, pekerjaan.id, pekerjaan.nama,
			penghasilan.id, penghasilan.keterangan, person.pembuat, person.wafat
			FROM person LEFT JOIN pendidikan ON person.id_pendidikan = pendidikan.id
			LEFT JOIN pekerjaan ON person.id_pekerjaan = pekerjaan.id
			LEFT JOIN penghasilan ON person.id_penghasilan = penghasilan.id `
	personKeluarga = `SELECT p.uuid, p.nik, p.nokk, p.nama, k.wali, s.keterangan_dasar, k.kepada_id_person
			FROM person p INNER JOIN keluarga k ON p.uuid = k.kepada_id_person
			INNER JOIN status_keluarga s ON k.id_status_keluarga = s.id
			WHERE k.dasar_id_person IN (?)`
	personSelectByUUID       = personSelectBase + `WHERE person.uuid = ?`
	PersonSelectLimit        = personSelectBase + `limit ?, ?`
	personSettingPekerjaan   = `select * from pekerjaan`
	personSettingPendidikan  = `select * from pendidikan`
	personSettingPenghasilan = `select * from penghasilan`
)

type mysqlConfig struct {
	sync.Mutex
	DB *sql.DB
}

type Repository interface {
	GetUser(ctx context.Context, username, password string) (*models.Admin, error)
	GetPerson(ctx context.Context, page, limit int) ([]*models.Person, error)
	GetSettingPekerjaan(ctx context.Context) ([]*models.Pekerjaan, error)
	GetSettingPendidikan(ctx context.Context) ([]*models.Pendidikan, error)
	GetSettingPenghasilan(ctx context.Context) ([]*models.Penghasilan, error)
	GetKeluarga(ctx context.Context, uuid string) ([]*models.Keluarga, error)
	GetPersonByUUID(ctx context.Context, uuid string) (*models.Person, error)
	InsertPerson(ctx context.Context, a *models.Person, pembuat string) error
	GetTotal(ctx context.Context) (*models.Total, error)
	GetTotalPerson(ctx context.Context) int

	AddPekerjaan(ctx context.Context, p *models.Pekerjaan) error
	AddPendidikan(ctx context.Context, p *models.Pendidikan) error
	AddPenghasilan(ctx context.Context, p *models.Penghasilan) error
}

func NewRepositoryConfig(DB *sql.DB) Repository {
	return &mysqlConfig{DB: DB}
}

func (m *mysqlConfig) GetUser(ctx context.Context, username, password string) (*models.Admin, error) {
	result := new(models.Admin)
	var sha = sha1.New()
	sha.Write([]byte(password))
	var encrypted = sha.Sum(nil)
	password = fmt.Sprintf("%x", encrypted)
	if err := m.DB.QueryRowContext(ctx, userLogin, username, password).Scan(
		&result.UUID, &result.Username, &result.Password, &result.Email,
		&result.Group, &result.Aktif,
	); err != nil {
		return nil, err
	}
	if result.Aktif == "T" {
		return nil, models.ErrAkunNoAktif
	}
	return result, nil
}

func (m *mysqlConfig) getPerson(ctx context.Context, query string, args ...interface{}) ([]*models.Person, error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Person, 0)
	for rows.Next() {
		P := new(models.Person)
		pekerjaan := new(models.Pekerjaan)
		pendidikan := new(models.Pendidikan)
		penghasilan := new(models.Penghasilan)
		if err := rows.Scan(
			&P.UUID, &P.NoKK, &P.NIK, &P.Nama, &P.JK, &P.TempatLahir, &P.TanggalLahir,
			&P.AnakKe, &P.JumSaudara, &P.Phone1, &P.Phone2, &pendidikan.Id, &pendidikan.Nama,
			&pekerjaan.Id, &pekerjaan.Nama,
			&penghasilan.Id, &penghasilan.Keterangan, &P.Pembuat, &P.Wafat,
		); err != nil {
			return nil, err
		}
		P.Umur = cekUmur(P.TanggalLahir)
		P.Pendidikan = pendidikan
		P.Pekerjaan = pekerjaan
		P.Penghasilan = penghasilan
		result = append(result, P)
	}
	return result, nil
}

func (m *mysqlConfig) GetPerson(ctx context.Context, page, limit int) ([]*models.Person, error) {
	if page > 1 {
		page = (page - 1) * limit
	}

	res, err := m.getPerson(ctx, PersonSelectLimit, page, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *mysqlConfig) GetPersonByUUID(ctx context.Context, uuid string) (*models.Person, error) {
	person, err := m.getPerson(ctx, personSelectByUUID, uuid)
	if err != nil {
		return nil, err
	}
	p := new(models.Person)
	if len(person) > 0 {
		p = person[0]
	} else {
		return nil, models.ErrNotFound
	}
	return p, nil
}

func (m *mysqlConfig) GetSettingPekerjaan(ctx context.Context) ([]*models.Pekerjaan, error) {
	rows, err := m.DB.QueryContext(ctx, personSettingPekerjaan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*models.Pekerjaan, 0)
	for rows.Next() {
		P := new(models.Pekerjaan)
		if err := rows.Scan(
			&P.Id, &P.Nama,
		); err != nil {
			return nil, err
		}

		result = append(result, P)
	}

	return result, nil
}

func (m *mysqlConfig) GetSettingPendidikan(ctx context.Context) ([]*models.Pendidikan, error) {
	rows, err := m.DB.QueryContext(ctx, personSettingPendidikan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*models.Pendidikan, 0)
	for rows.Next() {
		P := new(models.Pendidikan)
		if err := rows.Scan(
			&P.Id, &P.Nama,
		); err != nil {
			return nil, err
		}

		result = append(result, P)
	}

	return result, nil
}

func (m *mysqlConfig) GetSettingPenghasilan(ctx context.Context) ([]*models.Penghasilan, error) {
	rows, err := m.DB.QueryContext(ctx, personSettingPenghasilan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*models.Penghasilan, 0)
	for rows.Next() {
		P := new(models.Penghasilan)
		if err := rows.Scan(&P.Id, &P.Keterangan); err != nil {
			return nil, err
		}
		result = append(result, P)
	}
	return result, nil
}

func (m *mysqlConfig) GetKeluarga(ctx context.Context, uuid string) ([]*models.Keluarga, error) {
	rows, err := m.DB.QueryContext(ctx, personKeluarga, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*models.Keluarga, 0)

	for rows.Next() {
		K := new(models.Keluarga)
		if err := rows.Scan(
			&K.UUID, &K.NIK, &K.NoKK, &K.Nama, &K.SebagaiWali, &K.StatusRelasi, &K.UUIDPersonLawan,
		); err != nil {
			return nil, err
		}
		result = append(result, K)
	}

	return result, nil
}

func (m *mysqlConfig) InsertPerson(ctx context.Context, a *models.Person, pembuat string) error {
	m.Lock()
	defer m.Unlock()
	stmt, err := m.DB.PrepareContext(ctx, personInsert)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(
		ctx, uuid.Must(uuid.NewV4()), a.NoKK, a.NIK, a.Nama, a.JK, a.TempatLahir, a.TanggalLahir,
		a.AnakKe, a.JumSaudara, a.Phone1, a.Phone2, a.Pendidikan.Id, a.Pekerjaan.Id, a.Penghasilan.Id,
		pembuat,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		err := fmt.Errorf("Weird Behavior. Total affected : %d ", rowsAffected)
		return err
	}
	return nil
}

func cekUmur(kelahiran string) int {
	var format = "2006-01-02"
	t, err := time.Parse(format, kelahiran)
	if err != nil {
		return 0
	}
	s := time.Now()
	sekarang := s.Year()

	if t.Month() > s.Month() {
		sekarang-- // = sekarang - 1
	} else if t.Month() == s.Month() && t.Day() > s.Day() {
		sekarang--
	}
	return sekarang - t.Year()
}
