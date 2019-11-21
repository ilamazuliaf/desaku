package models

import "errors"

var (
	ErrInternalServerError = errors.New("Ops, Ada Masalah Pada Server")
	ErrNotFound            = errors.New("Ops, Url Tidak Ditemukan")
	ErrConflict            = errors.New("Ops, Data Sudah Ada")
	ErrBadRequest          = errors.New("Ops, Parameter Tidak Valid")
	ErrAkunNoAktif         = errors.New("Ops. Akun anda sudah di non aktifkan, silahkan hubungi Administrator")
)
