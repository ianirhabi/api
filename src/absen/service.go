package absen

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

type absen struct {
	absen *model.Absen `json:"absen"`
}

type Respon struct {
	Data interface{} `json:"data"`
}

var b Respon

func Getabsen(id int) (c interface{}, e error) {
	o := orm.NewOrm()
	var cc []*model.Absen
	o.Raw("select * from absen where id_user = ?", id).QueryRows(&cc)

	b.Data = cc
	return b, e
}

func Req(d Request) (data interface{}, e error) {

	o := orm.NewOrm()
	abs := model.Absen{Tanggal: d.Tanggal,
		Waktu:   d.Waktu,
		Hadir:   d.Kehadiran,
		Hari:    d.Hari,
		ID_USER: d.Iduser,
		Lat:     d.Lat,
		Long:    d.Long,
		Usr:     d.User}
	// ID_USER: d.Iduser}
	fmt.Println("debug kah ", &abs)
	b.Data = &abs
	o.Insert(b.Data)
	return b.Data, e
}
