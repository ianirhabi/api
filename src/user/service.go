package user

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

func Getuser(id int, uon model.User) (cc model.User, e error) {
	o := orm.NewOrm()
	if d := o.Raw("select * from user where id = ?", id).QueryRow(&uon); d == nil {
		fmt.Println("usernya adalah ", uon)
		update := orm.NewOrm()
		var cc []*model.User
		update.Raw("update user set usergroup= '1' where id = ?", id).QueryRow(&cc)
	}
	return uon, e
}
