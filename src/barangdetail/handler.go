package barangdetail

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Getbd(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	page := c.Param("page")
	code_category, _ := strconv.Atoi(c.Param("code_category"))
	if data, m := Getitemdetail(usergrup, page, code_category); m == nil {
		return c.JSON(http.StatusOK, &data)
	}
	return e
}

func PostDetailItem(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	var r RequestBarangDetail
	if err := c.Bind(&r); err == nil {
		if data, m := PostItemDetail(r, usergrup); m == nil {
			return c.JSON(http.StatusOK, &data)
		}
	}
	return e
}
