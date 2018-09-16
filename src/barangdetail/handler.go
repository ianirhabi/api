package barangdetail

import (
	"fmt"
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

func DeleteDetailItem(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		if data, err := DeleteData(usergrup, int64(id)); err == nil {
			return c.JSON(http.StatusOK, &data)
		}
	} else {
		var stat = err
		fmt.Println("status", usergrup, id, err)
		return c.JSON(http.StatusNotFound, &stat)
	}

	return
}

func UpdateDetailItem(c echo.Context) (e error) {
	usergrup := c.Param("usergrup")
	var r UpdateBarangDetail
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		if err := c.Bind(&r); err == nil {
			if data, err := UpdateDetail(usergrup, int64(id), r); err == nil {
				return c.JSON(http.StatusOK, &data)
			}
		} else {
			fmt.Println(err)
			return c.JSON(http.StatusNotFound, err)
		}
	} else {
		var stat = err
		fmt.Println("status", usergrup, id, err)
		return c.JSON(http.StatusNotFound, &stat)
	}
	return
}
