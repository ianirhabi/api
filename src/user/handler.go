package user

import (
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
	//id, _ := strconv.Atoi(c.Param("id"))
	//delete(users, id)
	return c.NoContent(http.StatusNoContent)
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
