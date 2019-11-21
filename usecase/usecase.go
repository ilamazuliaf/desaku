package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ilamazuliaf/desaku/models"
	"github.com/ilamazuliaf/desaku/repository"
)

type usecaseConfig struct {
	mysql          repository.Repository
	contextTimeout time.Duration
}

type Usecase interface {
	GetUser(c context.Context, username, password string) (*models.Admin, error)
	GetPerson(c context.Context, page, limit int) ([]*models.Person, int, error)
	SettingPekerjaan(c context.Context) ([]*models.Pekerjaan, error)
	SettingPendidikan(c context.Context) ([]*models.Pendidikan, error)
	SettingPenghasilan(c context.Context) ([]*models.Penghasilan, error)
	GetDetailPerson(c context.Context, uuid string) (*models.DetailPerson, error)
	InsertPerson(c context.Context, m *models.Person, pembuat string) error
	GetTotal(c context.Context) (*models.Total, error)
	Persentase(c context.Context) (*models.Persentase, error)

	AddPekerjaan(c context.Context, p *models.Pekerjaan) error
	AddPendidikan(c context.Context, p *models.Pendidikan) error
	AddPenghasilan(c context.Context, p *models.Penghasilan) error
}

func NewUsecaseConfig(mysql repository.Repository, contextTimeout time.Duration) Usecase {
	return &usecaseConfig{mysql, contextTimeout}
}

func (d *usecaseConfig) GetUser(c context.Context, username, password string) (*models.Admin, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	user, err := d.mysql.GetUser(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *usecaseConfig) GetTotal(c context.Context) (*models.Total, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	return d.mysql.GetTotal(ctx)
}

func (d *usecaseConfig) GetPerson(c context.Context, page, limit int) ([]*models.Person, int, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	Persons, err := d.mysql.GetPerson(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}
	// totalPerson := d.mysql.GetTotalPerson(ctx)
	totalPerson := d.mysql.GetTotalPerson(ctx)
	// if err != nil {
	// 	return nil, nil, err
	// }
	return Persons, totalPerson, nil
}

func (d *usecaseConfig) SettingPendidikan(c context.Context) ([]*models.Pendidikan, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	pendidikan, err := d.mysql.GetSettingPendidikan(ctx)
	if err != nil {
		return nil, err
	}
	return pendidikan, nil
}

func (d *usecaseConfig) SettingPekerjaan(c context.Context) ([]*models.Pekerjaan, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	pekerjaan, err := d.mysql.GetSettingPekerjaan(ctx)
	if err != nil {
		return nil, err
	}
	return pekerjaan, nil
}

func (d *usecaseConfig) SettingPenghasilan(c context.Context) ([]*models.Penghasilan, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	penghasilan, err := d.mysql.GetSettingPenghasilan(ctx)
	if err != nil {
		return nil, err
	}
	return penghasilan, nil
}

func (d *usecaseConfig) GetDetailPerson(c context.Context, uuid string) (*models.DetailPerson, error) {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	result := new(models.DetailPerson)
	person, err := d.mysql.GetPersonByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	keluarga, err := d.mysql.GetKeluarga(ctx, uuid)
	if err != nil {
		return nil, err
	}
	result.Person = person
	result.Keluarga = keluarga
	return result, nil
}

func (d *usecaseConfig) InsertPerson(c context.Context, m *models.Person, pembuat string) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()

	if err := d.mysql.InsertPerson(ctx, m, pembuat); err != nil {
		return err
	}
	return nil
}

func (d *usecaseConfig) AddPekerjaan(c context.Context, p *models.Pekerjaan) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	if err := d.mysql.AddPekerjaan(ctx, p); err != nil {
		return err
	}
	return nil
}

func (d *usecaseConfig) AddPendidikan(c context.Context, p *models.Pendidikan) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	if err := d.mysql.AddPendidikan(ctx, p); err != nil {
		return err
	}
	return nil
}
func (d *usecaseConfig) AddPenghasilan(c context.Context, p *models.Penghasilan) error {
	ctx, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	if err := d.mysql.AddPenghasilan(ctx, p); err != nil {
		return err
	}
	return nil
}

func (h *usecaseConfig) Persentase(c context.Context) (*models.Persentase, error) {
	p := new(models.Persentase)
	ctx, cancel := context.WithTimeout(c, h.contextTimeout)
	defer cancel()
	total, err := h.GetTotal(ctx)
	if err != nil {
		return nil, err
	}
	p.JK = &models.PersentaseJK{
		Laki:      fmt.Sprintf("%.02f", (float32(total.Laki)/float32(total.Person))*100),
		Perempuan: fmt.Sprintf("%.02f", (float32(total.Perempuan)/float32(total.Person))*100),
	}
	p.Pendidikan = &models.PersentasePendidikan{
		SD:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.SD)/float64(total.Person))*100),
		SLTP: fmt.Sprintf("%.02f", (float64(total.Pendidikan.SLTP)/float64(total.Person))*100),
		SLTA: fmt.Sprintf("%.02f", (float64(total.Pendidikan.SLTA)/float64(total.Person))*100),
		D1:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.D1)/float64(total.Person))*100),
		D2:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.D2)/float64(total.Person))*100),
		D3:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.D3)/float64(total.Person))*100),
		S1:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.S1)/float64(total.Person))*100),
		S2:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.S2)/float64(total.Person))*100),
		S3:   fmt.Sprintf("%.02f", (float64(total.Pendidikan.S3)/float64(total.Person))*100),
	}

	return p, nil
}
