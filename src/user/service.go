package user

import (
	"fmt"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"
)

func Getuser(id int, uon model.User) (e error) {
	o := orm.NewOrm()
	fmt.Println("masuk kesini ", uon)
	if d := o.Raw("select * from user where id = ?", id).QueryRow(&uon); d == nil {
		fmt.Println("usernya adalah ", uon)
	}
	return e
}
