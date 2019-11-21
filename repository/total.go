package repository

import (
	"context"

	"github.com/ilamazuliaf/desaku/models"
)

const (
	totalBase      = `select count(*) from person `
	totalLaki      = totalBase + `where jk="L"`
	totalPerempuan = totalBase + `where jk="P"`
	totalSD        = totalBase + `where id_pendidikan = 1`
	totalSLTP      = totalBase + `where id_pendidikan = 2`
	totalSLTA      = totalBase + `where id_pendidikan = 3`
	totalD1        = totalBase + `where id_pendidikan = 4`
	totalD2        = totalBase + `where id_pendidikan = 5`
	totalD3        = totalBase + `where id_pendidikan = 6`
	totalS1        = totalBase + `where id_pendidikan = 7`
	totalS2        = totalBase + `where id_pendidikan = 8`
	totalS3        = totalBase + `where id_pendidikan = 9`
)

func (m *mysqlConfig) GetTotalPerson(ctx context.Context) int {
	var total int = 0
	if err := m.DB.QueryRowContext(ctx, totalBase).Scan(&total); err != nil {
		return total
	}
	return total
}

func (m *mysqlConfig) GetTotal(ctx context.Context) (*models.Total, error) {
	var err error
	p := new(models.Total)
	pendidikan := new(models.TotalPendidikan)
	// if err = m.DB.QueryRowContext(ctx, totalBase).Scan(&p.Person); err != nil {
	// 	return nil, err
	// }
	p.Person = m.GetTotalPerson(ctx)

	if err = m.DB.QueryRowContext(ctx, totalPerempuan).Scan(&p.Perempuan); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalLaki).Scan(&p.Laki); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalSD).Scan(&pendidikan.SD); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalSLTP).Scan(&pendidikan.SLTP); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalSLTA).Scan(&pendidikan.SLTA); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalD1).Scan(&pendidikan.D1); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalD2).Scan(&pendidikan.D2); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalD3).Scan(&pendidikan.D3); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalS1).Scan(&pendidikan.S1); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalS2).Scan(&pendidikan.S2); err != nil {
		return nil, err
	}

	if err = m.DB.QueryRowContext(ctx, totalS3).Scan(&pendidikan.S3); err != nil {
		return nil, err
	}

	p.Pendidikan = pendidikan
	return p, nil
}
