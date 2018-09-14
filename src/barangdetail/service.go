package barangdetail

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

func Getitemdetail(usrgrup string, page string, kt int) (d interface{}, e error) {
	var barangdetail []*model.Item_barang_detail
	o := orm.NewOrm()
	var b Respons
	if usrgrup == "1" || usrgrup == "2" {
		if d, x := o.Raw("SELECT * FROM item_barang_detail  where code_category =? order by id ASC limit "+page, kt).QueryRows(&barangdetail); x == nil {
			fmt.Println("debug berhasil ==== ", d, x)
			b.Status = "berhasil"
			b.Data = barangdetail
			b.Total = page
		} else {
			fmt.Println("debug errornya ==== ", d, x)
			b.Status = "gagal"
		}
	}
	return b, e
}

func PostItemDetail(r RequestBarangDetail, usergrup string) (d interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()
	var validasi []*model.Item_barang_detail
	if usergrup == "1" || usergrup == "2" {
		if d, x := o.Raw("select * from item_barang_detail where code =?", r.Code).QueryRows(&validasi); x == nil {
			if d == 0 {
				input := model.Item_barang_detail{Code: r.Code,
					Stock:               r.Stock,
					Hargapokokpenjualan: r.Hp,
					Hargajual:           r.Hj,
					Codecategory:        r.Code_kategory,
					Deskripsi:           r.Name,
					Created:             r.Created,
					UpdateDate:          r.Created,
					UserId:              r.UserID,
				}
				if df, e := o.Insert(&input); e == nil {
					Res.Status = "sukses"
					Res.Data = input
					fmt.Println("debug kirim barang sukses === ", df, e)
				} else {
					fmt.Println("debug kirim barang gagal === ", r.Code_kategory, df, e)
					Res.Status = "gagal"
				}
			} else {
				Res.Status = "false"
				Res.Data = d
			}
		}
	} else {
		Res.Status = "Anda Tidak di izinkan akses disini"
	}
	return Res, e
}
