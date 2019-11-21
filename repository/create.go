package repository

import (
	"context"

	"github.com/ilamazuliaf/desaku/models"
)

const (
	addPekerjaan   = `insert pekerjaan set nama = ?`
	addPendidikan  = `insert pendidikan set nama = ?`
	addPenghasilan = `insert penghasilan set keterangan = ?`
)

func (m *mysqlConfig) AddPekerjaan(ctx context.Context, p *models.Pekerjaan) error {
	stmt, err := m.DB.PrepareContext(ctx, addPekerjaan)
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, p.Nama); err != nil {
		return err
	}

	return nil
}

func (m *mysqlConfig) AddPendidikan(ctx context.Context, p *models.Pendidikan) error {
	stmt, err := m.DB.PrepareContext(ctx, addPendidikan)
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, p.Nama); err != nil {
		return err
	}
	return nil
}

func (m *mysqlConfig) AddPenghasilan(ctx context.Context, p *models.Penghasilan) error {
	stmt, err := m.DB.PrepareContext(ctx, addPenghasilan)
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, p.Keterangan); err != nil {
		return err
	}
	return nil
}
