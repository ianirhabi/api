package absen

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) URLMapping(r *echo.Group) {
	r.GET("/:id", h.absen)
	r.POST("", h.kirimabsen)
	r.GET("/:id/:from/:to", h.getabsendate)
}

func (h *Handler) absen(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	if cc, m := Getabsen(id); m == nil {
		return c.JSON(http.StatusOK, cc)
		// return c.JSON(http.StatusOK, &cc)
	}
	return e
}

func (h *Handler) getabsendate(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	from := c.Param("from")
	to := c.Param("to")
	fmt.Println("debug id", id)
	fmt.Println("debug from", from)
	fmt.Println("debug to", to)
	if cc, m := Getabsenwithdate(from, to, id); m == nil {
		return c.JSON(http.StatusOK, cc)
		// return c.JSON(http.StatusOK, &cc)
	}
	return e
}

func (h *Handler) kirimabsen(c echo.Context) (e error) {
	var r Request
	if err := c.Bind(&r); err == nil {
		if data, e := Req(r); e == nil {
			return c.JSON(http.StatusCreated, data)
		}
	}
	return e
}
