package absen

import (
	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

type absen struct {
	absen *model.Absen `json:"absen"`
}

func Getabsen(id int) (cc []*model.Absen, e error) {
	o := orm.NewOrm()
	o.Raw("select * from absen where id_user = ?", id).QueryRows(&cc)

	return cc, e
}
