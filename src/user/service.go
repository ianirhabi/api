package user

import (
	"fmt"
	"time"

	"github.com/alfatih/beego/orm"
	jwt "github.com/dgrijalva/jwt-go"
	"retrobarbershop.com/retro/api/model"
)

func Getuser(id int, uon model.User) (cc model.User, e error) {
	o := orm.NewOrm()
	if d := o.Raw("select * from user where id = ?", id).QueryRow(&uon); d == nil {
		update := orm.NewOrm()
		var cc []*model.User
		update.Raw("update user set notifikasi= '0' where id = ?", id).QueryRow(&cc)
	}
	return uon, e
}

func Getalluser(cc []*model.User) (d interface{}, e error) {
	o := orm.NewOrm()
	var res ResponsAlluser
	if qs := o.QueryTable("user"); qs != nil { // mengecek tabel user apakah ada atau tidak... aklau tidak ada akan terjadi panic
		if _, x := o.Raw("select * from user").QueryRows(&cc); x == nil {
			res.Status = "berhasil"
			res.Data = cc
		} else {
			res.Status = "gagal"
		}
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
		a.Usr.Notifikasi = row.Notifikasi
		a.Usr.Usergrup = row.Usergrup
		a.Usr.Ianmonitor = row.Ianmonitor
		a.Usr.Name = row.Name
		a.Usr.NamaFoto = row.NamaFoto
		if row.User == a.Name && row.Pass == a.Passw {
			data.Status = "sukses"
			data.Jam = time.Now()
			data.Data = a

			// Set claims
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["username"] = "Jon Snow"
			claims["password"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
			data.Token = t

			if err != nil {
				return data, err
			}
			break
		} else {
			data.Status = "gagal"
		}
	}
	return data, e
}

func Updateuser(Req UpdateRequest, id int) (d interface{}, e error) {
	o := orm.NewOrm()
	var res Respons
	fmt.Println("idnya ", id)
	var cc []*model.User
	fmt.Println("password baru ", Req.Password)
	if x := o.Raw("update user set password_retro =? where id = ? ", Req.Password, id).QueryRow(&cc); x != nil {
		res.Status = "sukses"
		res.Data = cc
	}
	return res, e
}

func Deleteuser(id int) (d interface{}, e error) {
	o := orm.NewOrm()
	var res Respons
	var uon model.User
	if d := o.Raw("SELECT * FROM user where id=?", id).QueryRow(&uon); d == nil {
		o.Raw("delete from absen where id_user = ? ", id).QueryRow(&uon)
		if x := o.Raw("delete from user where id = ? ", id).QueryRow(&uon); x != nil {
			res.Status = "sukses"
			res.Data = "Anda Telah Berhasil Delete User"
		} else {
			fmt.Println("Anda gagal delete user")
		}
	} else {
		res.Status = "gagal"
		res.Data = "Anda gagal delete User"
	}
	return res, e
}

func Add(r AddRequest) (d interface{}, e error) {
	var res Respons
	o := orm.NewOrm()
	abs := model.User{
		Notifikasi: "",
		Name:       r.Name,
		User:       r.User,
		Pass:       r.Pass,
		NamaFoto:   "default.jpg",
		Latitude:   "",
		Longtitude: "",
		Usergrup:   r.Usergrup,
		Ianmonitor: "1"}
	// ID_USER: d.Iduser}
	fmt.Println("debug kah ", &abs)
	if id, x := o.Insert(&abs); x == nil {
		res.Data = &abs
		fmt.Println("debug ", id, x)
		res.Status = "berhasil"
	} else {
		res.Status = "gagal"
		fmt.Println("debug ", id, x)
	}
	return res, e
}
