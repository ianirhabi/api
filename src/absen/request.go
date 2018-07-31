package absen

type Request struct {
	Tanggal   string `json:"Tanggal"`
	Waktu     string `json:"Waktu"`
	Kehadiran string `json:"Hadir"`
	Iduser    int64  `json:"ID_USER"`
	Hari      string `json:"Hari"`
	Lat       string `json:"Lat"`
	Long      string `json:"Long"`
	User      string `json:"user"`
	Alasan    string `json:"alasan"`
}

type Respon struct {
	Res  string      `json:"res"`
	Data interface{} `json:"data"`
}
