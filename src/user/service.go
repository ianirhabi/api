package user

import (
	"fmt"
	"net/http"

	"github.com/alfatih/beego/orm"
	"github.com/labstack/echo"
	"retrobarbershop.com/retro/api/model"
)

func Getuser(id int, uon model.User) (e error) {
	var c echo.Context
	o := orm.NewOrm()
	fmt.Println("masuk kesini ", uon)
	if d := o.Raw("select * from user where id = ?", id).QueryRow(&uon); d == nil {
		fmt.Println("usernya adalah ", uon)
		return c.JSON(http.StatusOK, &uon)
	}
	return e
}
