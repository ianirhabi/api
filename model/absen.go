package model

import (
	"github.com/alfatih/beego/orm"
)

func init() {
	orm.RegisterModel(new(Absen))
}

// Model Struct
type Absen struct {
	Id      int    `orm:"column(id);auto"json:"-"`
	Tanggal string `orm:"column(tanggal);size(100)"`
	Waktu   string `orm:"column(waktu);size(100)"`
	Hadir   string `orm:"column(kehadiran);size(100)"`
	ID_USER int64  `orm:"column(id_user);size(100)"`
	Hari    string `orm:"column(hari);size(100)"`
	Lat     string `orm:"column(lat);size(100)"`
	Long    string `orm:"column(long);size(100)"`
	Usr     string `orm:"column(user);size(100)"`

	//User *User `orm:"column(id_user);rel(fk)"json:"user,omitempty"`
}
