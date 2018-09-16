package barangdetail

import (
	"fmt"
	"strconv"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

func Getitemdetail(usrgrup string, page string, kt int) (d interface{}, e error) {
	var barangdetail []*model.Item_barang_detail
	o := orm.NewOrm()
	var b Respons
	if usrgrup == "1" || usrgrup == "2" {
		if d, x := o.Raw("SELECT * FROM item_barang_detail  where code_category =? order by id DESC limit "+page, kt).QueryRows(&barangdetail); x == nil {
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

func DeleteData(usergrup string, id int64) (d interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()

	if usergrup == "1" || usergrup == "2" {
		if status, err := o.Delete(&model.Item_barang_detail{Id: id}); err == nil {
			Res.Status = "sukses"
			Res.Data = "berhasil menghapus data"
			fmt.Println("debug berhasil detailitem : ", status, nil)
		} else {
			Res.Status = "gagal"
			fmt.Println("debug gagal detailitem : ", status, nil)
		}
	}
	return Res, e
}

func UpdateDetail(usergrup string, id int64, r UpdateBarangDetail) (d interface{}, e error) {
	var Res Respons
	o := orm.NewOrm()
	var updatedata []*model.Item_barang_detail
	if usergrup == "1" || usergrup == "2" {
		if d, x := o.Raw("select * from item_barang_detail where code =?", r.Code).QueryRows(&updatedata); x == nil {
			if d == 0 {
				datadetail := model.Item_barang_detail{Id: id}
				if o.Read(&datadetail) == nil {
					fmt.Println("debug berhasil update", d)
					datadetail.Code = r.Code
					datadetail.Codecategory = r.Code_kategory
					datadetail.UpdateDate = r.Created
					datadetail.Deskripsi = r.Name
					datadetail.Hargajual = r.Hj
					datadetail.Hargapokokpenjualan = r.Hp
					datadetail.Stock = r.Stock
					datadetail.UserId = r.UserID
					if stat, err := o.Update(&datadetail); err == nil {
						Res.Status = "sukses"
						Res.Data = datadetail
						//konversi tipe data int ke tipe data string
						var totalData = int64(d)
						var str = strconv.FormatInt(totalData, 10)
						Res.Total = str
					} else {
						fmt.Println("debug gagal update", stat, err)
						Res.Status = "gagal"
						Res.Data = stat
						//konversi tipe data int ke tipe data string
						var totalData = int64(d)
						var str = strconv.FormatInt(totalData, 10)
						Res.Total = str
					}
				}
			} else {
				Res.Status = "false"
				Res.Data = "kode sudah ada silahkan coba yang lain"
				fmt.Println("debug berhasil update", d)
				//konversi tipe data int ke tipe data string
				var totalData = int64(d)
				var str = strconv.FormatInt(totalData, 10)
				Res.Total = str
			}
		} else {
			fmt.Println("debug gagal : ", d, x)
		}
	} else {
		Res.Data = "forbidden 404"
	}
	return Res, e
}
