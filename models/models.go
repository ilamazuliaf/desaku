package models

import (
	"database/sql"
	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
	validator "gopkg.in/go-playground/validator.v9"
)

type MyClaims struct {
	jwt.StandardClaims
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Group    string `json:"group"`
}

type Admin struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group    string `json:"group"`
	Aktif    string `json:"aktif"`
}

type Person struct {
	UUID         string  `json:"uuid"`
	NoKK         string  `json:"nokk" validate:"required"`
	NIK          string  `json:"nik" validate:"required"`
	Nama         string  `json:"nama" validate:"required"`
	JK           string  `json:"jenis_kelamin" validate:"required"`
	TempatLahir  string  `json:"tempat_lahir" validate:"required"`
	TanggalLahir string  `json:"tanggal_lahir" validate:"required"`
	AnakKe       int     `json:"anak_ke" validate:"required,numeric,gte=0,lte=25"`
	JumSaudara   int     `json:"jum_saudara" validate:"required,numeric,gte=0,lte=25"`
	Phone1       *string `json:"phone1"`
	Phone2       *string `json:"phone2"`
	Umur         int     `json:"umur"`
	// Pendidikan   *string `json:"pendidikan"`
	// Pekerjaan    *string `json:"pekerjaan"`
	// Penghasilan  *string `json:"penghasilan"`
	// Pendidikan  Pendidikan  `json:"pendidikan,omitempty"`
	// Pekerjaan   Pekerjaan   `json:"pekerjaan"`
	// Penghasilan Penghasilan `json:"penghasilan"`
	*Pendidikan
	*Pekerjaan
	*Penghasilan
	Pembuat string `json:"created_by"`
	Wafat   string `json:"wafat"`
}

type Pendidikan struct {
	Id   *int    `json:"id_pendidikan,omitempty"`
	Nama *string `json:"pendidikan" validate:"required"`
}

type Pekerjaan struct {
	Id   *int    `json:"id_pekerjaan,omitempty"`
	Nama *string `json:"pekerjaan" validate:"required"`
}

type Penghasilan struct {
	Id         *int    `json:"id_penghasilan,omitempty"`
	Keterangan *string `json:"penghasilan" validate:"required"`
}

type Keluarga struct {
	UUID            string `json:"uuid"`
	NIK             string `json:"nik"`
	NoKK            string `json:"nokk"`
	Nama            string `json:"nama"`
	StatusRelasi    string `json:"status_relasi"`
	SebagaiWali     string `json:"wali"`
	UUIDPersonLawan string `json:"uuid_person_lawan"`
}

type DetailPerson struct {
	*Person
	Keluarga []*Keluarga `json:"keluarga,omitempty"`
}

type NullString struct {
	sql.NullString
}

type NullInt struct {
	sql.NullInt64
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns NullInt) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Int64)
}

// Validasi struct by tag
type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
