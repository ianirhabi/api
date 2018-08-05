package model

import "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(Item_barang_detail))
}

type Item_barang_detail struct {
	Id                  int    `orm:"column(id);auto"json:"id"`
	Code                string `orm:"column(code);size(100);null"json:"code"`
	Stock               string `orm:"column(stock);size(100);null"json:"stock"`
	Hargapokokpenjualan int64  `orm:"column(harga_pokok_penjualan);size(100);null"json:"harga_pokok"`
	Hargajual           int64  `orm:"column(harga_jual);size(100);null"json:"harga_jual"`
	Codecategory        string `orm:"column(code_category);size(100);null"json:"harga_jual"`
	Deskripsi           string `orm:"column(deskripsi);size(100);null"json:"deskripsi"`
}
