package model

// Model Struct
type User struct {
	Id         int    `orm:"column(id);auto"`
	Usergup    string `orm:"column(usergroup);size(100)"`
	Name       string `orm:"column(nama);size(100)"`
	User       string `orm:"column(user_retro);size(100)"`
	Pass       string `orm:"column(password_retro);size(100)"`
	NamaFoto   string `orm:"column(nama_foto);size(100)"`
	Latitude   string `orm:"column(latitude);size(100)"`
	Longtitude string `orm:"column(longtitude);size(100)"`
}
