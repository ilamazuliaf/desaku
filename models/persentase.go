package models

type Persentase struct {
	Pendidikan *PersentasePendidikan `json:"pendidikan"`
	JK         *PersentaseJK         `json:"jenis_kelamin"`
}

type PersentasePendidikan struct {
	SD   string `json:"sd"`
	SLTP string `json:"sltp"`
	SLTA string `json:"slta"`
	D1   string `json:"d1"`
	D2   string `json:"d2"`
	D3   string `json:"d3"`
	S1   string `json:"s1"`
	S2   string `json:"s2"`
	S3   string `json:"s3"`
}

type PersentaseJK struct {
	Laki      string `json:"laki-laki"`
	Perempuan string `json:"perempuan"`
}
