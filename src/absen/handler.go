package absen

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Absendetail(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	if cc, m := Getabsen(id); m == nil {
		return c.JSON(http.StatusOK, cc)
		// return c.JSON(http.StatusOK, &cc)
	}
	return e
}

func Getabsendate(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	from := c.Param("from")
	to := c.Param("to")
	if cc, m := Getabsenwithdate(from, to, id); m == nil {
		return c.JSON(http.StatusOK, cc)
		// return c.JSON(http.StatusOK, &cc)
	}
	return e
}

func Kirimabsen(c echo.Context) (e error) {
	var r Request
	if err := c.Bind(&r); err == nil {
		if data, e := Req(r); e == nil {
			return c.JSON(http.StatusCreated, data)
		}
	}
	return e
}

func Kirimzin(c echo.Context) (e error) {
	var r Request
	if err := c.Bind(&r); err == nil {
		if data, e := Izin(r); e == nil {
			return c.JSON(http.StatusOK, &data)
		}
	}

	return e
}
