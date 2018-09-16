package model

import "github.com/alfatih/beego/orm"

func init() {
	orm.RegisterModel(new(Item_barang))
}

type Item_barang struct {
	Id           int64  `orm:"column(id);auto"json:"id"`
	ItemCategory string `orm:"column(item_category);size(100);null"json:"item_category"`
	CodeItem     string `orm:"column(code_item);size(100);null"json:"code_item"`
	Created      string `orm:"column(created);size(100);null"json:"created"`
	UserId       int64  `orm:"column(user_id);size(100);null"json:"user_id"`
}
