package barang

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

func Inputitem(r Requestbarang, usergrup string, iduser int) (i interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()
	var validasi []*model.Item_barang
	//for n := 0; n < 360000000; n++ {
	//CI := strconv.Itoa(n) // konversi int ke string
	if usergrup == "1" || usergrup == "2" {
		if d, x := o.Raw("select * from item_barang where code_item =?", r.CodeITEM).QueryRows(&validasi); x == nil {
			if d == 0 {
				fmt.Println("debug exist ", d)
				input := model.Item_barang{ItemCategory: r.Itemcategory,
					CodeItem: r.CodeITEM,
					Created:  r.Created,
					UserId:   r.UserId,
				}

				if df, e := o.Insert(&input); e == nil {
					Res.Status = "sukses"
					Res.Data = input
					fmt.Println("debug kirim barang sukses === ", df, e)
				} else {
					fmt.Println("debug kirim barang gagal === ", df, e)
					Res.Status = "gagal"
				}
			} else {
				Res.Status = "false"
				Res.Data = "kode yang Anda masukan sudah ada, coba kode yang lain"
				fmt.Println("kode sudah ada ", d)
			}
		}
	} else {
		Res.Status = "Anda Tidak di izinkan akses disini"
	}
	//}
	return Res, e
}

func Getitem(a string, page string) (i interface{}, e error) {
	var barang []*model.Item_barang
	o := orm.NewOrm()
	var b Respons
	if a == "1" || a == "2" {
		if d, x := o.Raw("SELECT * FROM item_barang order by id DESC limit " + page).QueryRows(&barang); x == nil {
			var count int
			o.Raw("select count(*) as Count from item_barang").QueryRow(&count)
			fmt.Println("debug error ==== ", d, x)
			b.Status = "berhasil"
			b.Data = barang
			b.Total = page
		} else {
			fmt.Println("debug errornya ==== ", d, x)
			b.Status = "gagal"
		}
	} else {
		b.Status = "Anda dilarang mengambil data"
	}
	return b, e
}

func Updateitem(r Requestbarang, usergrup string, iduser int, idbarang int64) (i interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()
	var validasi []*model.Item_barang

	if usergrup == "1" || usergrup == "2" {
		if d, x := o.Raw("select * from item_barang where code_item =? and item_category =?", r.CodeITEM, r.Itemcategory).QueryRows(&validasi); x == nil {
			if d == 0 {
				barangitem := model.Item_barang{Id: idbarang}
				if o.Read(&barangitem) == nil {
					barangitem.ItemCategory = r.Itemcategory
					barangitem.CodeItem = r.CodeITEM
					barangitem.Created = r.Created
					barangitem.UserId = r.UserId
					if _, err := o.Update(&barangitem); err == nil {
						Res.Data = barangitem
						Res.Status = "berhasil"
					} else {
						Res.Status = "gagal"
					}
				} else {
					Res.Status = "Tidak Ada di daftar list"
				}
			} else {
				Res.Status = "false"
				Res.Data = "kode yang Anda masukan sudah ada, coba kode yang lain"
				fmt.Println("kode sudah ada ", d)
			}
		} else {
			Res.Status = "gagal"
			Res.Data = "somthing wrong"
		}

	} else {
		Res.Status = "Anda Tidak di izinkan akses disini"
	}
	return Res, e
}

func Deleteitem(usergrup string, idbarang int64) (i interface{}, e error) {
	o := orm.NewOrm()
	var Res Respons
	var delete []*model.Item_barang
	if usergrup == "1" || usergrup == "2" {
		if f, status := o.Raw("delete from item_barang_detail where code_category=?", idbarang).QueryRows(&delete); status == nil {
			if g, err := o.Delete(&model.Item_barang{Id: idbarang}); err == nil {
				Res.Status = "sukses"
				Res.Data = "berhasil menghapus data"
			} else {
				Res.Status = "gagal"
				fmt.Println("debug gagal dua ", g)
			}
		} else {
			Res.Status = "gagal"
			fmt.Println("debug gagal ", f)
		}
	} else {
		Res.Status = "Anda Tidak di izinkan"
	}
	return Res, e
}
