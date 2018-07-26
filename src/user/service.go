package user

import (
	"fmt"
	"time"

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

func Getalluser(cc []*model.User) (d interface{}, e error) {
	o := orm.NewOrm()
	var res ResponsAlluser
	if _, x := o.Raw("select * from user").QueryRows(&cc); x == nil {
		res.Status = "berhasil"
		res.Data = cc
	}
	return res, e
}

func PostLogin(a Requestlogin) (res interface{}, e error) {
	o := orm.NewOrm()
	var data Respons
	//user := model.User{User: b.Name,
	//	Pass: b.Passw}
	// insert
	//data, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", data, err)
	var uon []*model.User
	o.Raw("SELECT * FROM user").QueryRows(&uon)
	for _, row := range uon {
		a.Usr.Id = row.Id
		a.Usr.User = row.User
		a.Usr.Pass = row.Pass
		a.Usr.Usergup = row.Usergup
		if row.User == a.Name && row.Pass == a.Passw {
			data.Status = "sukses"
			data.Jam = time.Now()
			data.Data = a
			break
		} else {
			data.Status = "gagal"
		}
	}
	return data, e
}
