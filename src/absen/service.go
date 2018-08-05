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
	if d, x := o.Raw("select * from absen where id_user = ?", id).QueryRows(&cc); x == nil {
		fmt.Println("errornya ==== ", d, x)
		b.Res = "berhasil"
		b.Data = cc

	} else {
		fmt.Println("errornya ==== ", d, x)
		b.Res = "gagal"
	}
	return b, e
}

func Getabsenwithdate(from string, to string, id int) (c interface{}, e error) {
	o := orm.NewOrm()
	var cc []*model.Absen
	if _, d := o.Raw("SELECT * FROM absen where tanggal between '"+from+"' AND '"+to+"' AND id_user = ? ", id).QueryRows(&cc); d == nil {
		fmt.Println("debug error === ", d)
		b.Res = "berhasil"
		b.Data = cc
	} else {
		fmt.Println("debug error === ", d)
		b.Res = "gagal"
	}
	return b, e
}

func Req(d Request) (data interface{}, e error) {
	o := orm.NewOrm()
	update := orm.NewOrm()
	var cc []*model.User
	update.Raw("update user set notifikasi= '2',latitude =?,longtitude=? where id = ?", d.Lat, d.Long, d.Iduser).QueryRows(&cc)
	update.Raw("update user set notifikasi= '2' where user_retro = 'superadmin'").QueryRows(&cc)
	abs := model.Absen{Tanggal: d.Tanggal,
		Waktu:   d.Waktu,
		Hadir:   d.Kehadiran,
		Hari:    d.Hari,
		ID_USER: d.Iduser,
		Lat:     d.Lat,
		Long:    d.Long,
		Usr:     d.User,
		Pesan:   d.Alasan}
	// ID_USER: d.Iduser}
	fmt.Println("debug kah ", &abs)
	if _, e := o.Insert(&abs); e == nil {
		b.Res = "berhasil"
		b.Data = &abs
	} else {
		b.Res = "gagal"
	}
	return b, e
}

func Izin(r Request) (data interface{}, e error) {
	var Res Respon
	o := orm.NewOrm()
	update := orm.NewOrm()
	var cc []*model.User
	update.Raw("update user set notifikasi= '2',latitude =?,longtitude=? where id = ?", r.Lat, r.Long, r.Iduser).QueryRows(&cc)
	update.Raw("update user set notifikasi= '2' where user_retro = 'superadmin'").QueryRows(&cc)
	izin := model.Absen{
		Tanggal: r.Tanggal,
		Waktu:   r.Waktu,
		Hadir:   r.Kehadiran,
		Hari:    r.Hari,
		ID_USER: r.Iduser,
		Lat:     r.Lat,
		Long:    r.Long,
		Usr:     r.User,
		Pesan:   r.Alasan}
	if df, e := o.Insert(&izin); e == nil {
		Res.Res = "berhasil"
		Res.Data = &izin
	} else {
		fmt.Println("debug kirim izin gagal === ", df, e)
		Res.Res = "gagal"
	}
	return Res, e
}
