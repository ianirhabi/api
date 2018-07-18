package model

import "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(Absen))
}

// Model Struct
type Absen struct {
	Id      int    `orm:"column(id);auto"`
	Tanggal string `orm:"column(tanggal);size(100)"`
	Waktu   string `orm:"column(waktu);size(100)"`
	Hadir   string `orm:"column(kehadiran);size(100)"`
	User    *User  `orm:"column(id_user);rel(fk)"json:"user,omitempty"`
}
