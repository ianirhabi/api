package user

import (
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"retrobarbershop.com/retro/api/model"
)

func Login(c echo.Context) (e error) {

	var r Requestlogin
	if err := c.Bind(&r); err == nil {
		if cc, m := PostLogin(r); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

func Del(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("idnya", id)
	if data, er := Deleteuser(id); er == nil {
		return c.JSON(http.StatusOK, &data)
	}
	return e
}

func GetUser(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var uon model.User
	if cc, m := Getuser(id, uon); m == nil {
		return c.JSON(http.StatusOK, &cc)
	}
	return e
}

func GetAllUser(c echo.Context) (e error) {

	var uon []*model.User
	if cc, m := Getalluser(uon); m == nil {
		return c.JSON(http.StatusOK, &cc)
	}
	return e
}
func UpdateUser(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var r UpdateRequest
	if err := c.Bind(&r); err == nil {
		if cc, m := Updateuser(r, id); m == nil {
			return c.JSON(http.StatusOK, &cc)
		}
	}
	return e
}

func Adduser(c echo.Context) (e error) {
	var r AddRequest
	if err := c.Bind(&r); err == nil {
		if cc, m := Add(r); m == nil {
			return c.JSON(http.StatusCreated, &cc)
		}
	}
	return e
}
