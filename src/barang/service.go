package barang

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

func Inputitem(r Requestbarang, usergrup string, iduser int) (i interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()

	if usergrup == "1" || usergrup == "2" {
		input := model.Item_barang{ItemCategory: r.Itemcategory,
			CodeItem: r.CodeITEM,
			Created:  r.Created,
			UserId:   r.UserId}

		if df, e := o.Insert(&input); e == nil {
			Res.Status = "sukses"
			Res.Data = input
			fmt.Println("debug kirim barang sukses === ", df, e)
		} else {
			fmt.Println("debug kirim barang gagal === ", df, e)
			Res.Status = "gagal"
		}
	} else {
		Res.Status = "Anda Tidak di izinkan akses disini"
	}
	return Res, e
}

func Getitem(a string) (i interface{}, e error) {
	var barang []*model.Item_barang
	o := orm.NewOrm()
	var b Respons
	if a == "1" || a == "2" {
		if d, x := o.Raw("select * from item_barang").QueryRows(&barang); x == nil {

			var count int
			o.Raw("select count(*) as Count from item_barang").QueryRow(&count)
			fmt.Println("debug error ==== ", d, x)
			b.Status = "berhasil"
			b.Data = barang
			b.Total = count

		} else {
			fmt.Println("debug errornya ==== ", d, x)
			b.Status = "gagal"
		}
	} else {
		b.Status = "Anda dilarang mengambil data"
	}
	return b, e
}

func Updateitem(r Requestbarang, usergrup string, iduser int, idbarang int) (i interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()
	if usergrup == "1" || usergrup == "2" {
		barangitem := model.Item_barang{Id: idbarang}
		if o.Read(&barangitem) == nil {
			barangitem.ItemCategory = r.Itemcategory
			barangitem.CodeItem = r.CodeITEM
			barangitem.Created = r.Created
			barangitem.UserId = r.UserId
			if num, err := o.Update(&barangitem); err == nil {
				Res.Data = barangitem
				Res.Status = "berhasil"
				Res.Total = int(num)
			} else {
				Res.Status = "gagal update"
			}
		} else {
			Res.Status = "Tidak Ada di daftar list"
		}
	} else {
		Res.Status = "Anda Tidak di izinkan akses disini"
	}
	return Res, e
}

func Deleteitem(usergrup string, idbarang int) (i interface{}, e error) {
	o := orm.NewOrm()
	var Res Respons

	if usergrup == "1" || usergrup == "2" {
		if num, err := o.Delete(&model.Item_barang{Id: idbarang}); err == nil {
			Res.Status = "sukses"
			Res.Total = int(num)
		} else {
			Res.Status = "gagal"
		}
	} else {
		Res.Status = "Anda Tidak di izinkan"
	}
	return Res, e
}
