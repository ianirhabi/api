package absen

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

var b Respon

func Getabsen(id int) (c interface{}, e error) {
	o := orm.NewOrm()
	var cc []*model.Absen
	o.Raw("select * from absen where id_user = ?", id).QueryRows(&cc)
	b.Data = cc
	return b, e
}

func Getabsenwithdate(from string, to string, id int) (c interface{}, e error) {
	o := orm.NewOrm()
	var cc []*model.Absen
	o.Raw("SELECT * FROM absen where tanggal between '"+from+"' AND '"+to+"' AND id_user = ? ", id).QueryRows(&cc)
	b.Res = "berhasil"
	b.Data = cc
	return b, e
}

func Req(d Request) (data interface{}, e error) {

	o := orm.NewOrm()
	update := orm.NewOrm()
	var cc []*model.User
	update.Raw("update user set usergroup= '2' where id = ?", d.Iduser).QueryRows(&cc)
	abs := model.Absen{Tanggal: d.Tanggal,
		Waktu:   d.Waktu,
		Hadir:   d.Kehadiran,
		Hari:    d.Hari,
		ID_USER: d.Iduser,
		Lat:     d.Lat,
		Long:    d.Long,
		Usr:     d.User,
		Notif:   "1"}
	// ID_USER: d.Iduser}
	fmt.Println("debug kah ", &abs)
	b.Res = "berhasil"
	b.Data = &abs

	o.Insert(b.Data)
	return b, e
}
