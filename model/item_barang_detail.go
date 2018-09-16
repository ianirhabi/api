package model

import "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(Item_barang_detail))
}

type Item_barang_detail struct {
	Id                  int64  `orm:"column(id);auto"json:"id"`
	Code                string `orm:"column(code);size(100);null"json:"code"`
	Stock               string `orm:"column(stock);size(100);null"json:"stock"`
	Hargapokokpenjualan int64  `orm:"column(harga_pokok_barang);size(100);null"json:"harga_pokok"`
	Hargajual           int64  `orm:"column(harga_jual);size(100);null"json:"harga_jual"`
	Codecategory        int64  `orm:"column(code_category);size(100);null"json:"category_code"`
	Deskripsi           string `orm:"column(deskripsi);size(100);null"json:"deskripsi"`
	Created             string `orm:"column(created);size(100);null"json:"created"`
	UpdateDate          string `orm:"column(update);size(100);null"json:"update"`
	UserId              int64  `orm:"column(user_id);size(100);null"json:"userid"`
}
