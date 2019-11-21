package models

type Total struct {
	Person     int              `json:"person"`
	Laki       int              `json:"lelaki"`
	Perempuan  int              `json:"perempuan"`
	Pendidikan *TotalPendidikan `json:"pendidikan"`
}

type TotalPendidikan struct {
	SD   int `json:"sd"`
	SLTP int `json:"sltp"`
	SLTA int `json:"slta"`
	D1   int `json:"d1"`
	D2   int `json:"d2"`
	D3   int `json:"d3"`
	S1   int `json:"s1"`
	S2   int `json:"s2"`
	S3   int `json:"s3"`
}
