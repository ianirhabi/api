package barang

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Getbarang(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	if data, m := Getitem(usergrup); m == nil {
		return c.JSON(http.StatusOK, &data)
	}
	return e
}

func Postbarang(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	iduser, _ := strconv.Atoi(c.Param("user_id"))
	var r Requestbarang
	if err := c.Bind(&r); err == nil {
		if data, e := Inputitem(r, usergrup, iduser); e == nil {
			return c.JSON(http.StatusCreated, &data)
		}
	}
	return e
}

func Updatebarang(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	iduser, _ := strconv.Atoi(c.Param("user_id"))
	idbarang, _ := strconv.Atoi(c.Param("idbarang"))
	var r Requestbarang
	if err := c.Bind(&r); err == nil {
		if data, e := Updateitem(r, usergrup, iduser, idbarang); e == nil {
			return c.JSON(http.StatusCreated, &data)
		}
	}
	return e
}

func Deletebarang(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	barangid, d := strconv.Atoi(c.Param("barangid"))

	fmt.Println("idnya ", barangid, "errornya ", d)
	//fmt.Println(usergrup)
	if data, e := Deleteitem(usergrup, barangid); e == nil {
		return c.JSON(http.StatusOK, &data)
	}
	return e
}
