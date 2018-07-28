package model

// Model Struct
type User struct {
	Id         int    `orm:"column(id);auto"`
	Notifikasi string `orm:"column(notifikasi);size(100)"`
	Name       string `orm:"column(nama);size(100)"`
	User       string `orm:"column(user_retro);size(100)"`
	Pass       string `orm:"column(password_retro);size(100)"`
	NamaFoto   string `orm:"column(nama_foto);size(100)"`
	Latitude   string `orm:"column(latitude);size(100)"`
	Longtitude string `orm:"column(longtitude);size(100)"`
	Usergrup   string `orm:"column(usergrup);size(100)"`
	Ianmonitor string `orm:"column(ianmonitor);size(100)"`
}
