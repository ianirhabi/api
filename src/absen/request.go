package absen

type Request struct {
	Tanggal   string `json:"tanggal"`
	Waktu     string `json:"waktu"`
	Kehadiran string `json:"kehadiran"`
	Iduser    int64  `json:"id_user"`
	Hari      string `json:"hari"`
	Lat       string `json:"latitude"`
	Long      string `json:"longtitude"`
}
