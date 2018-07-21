package absen

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) URLMapping(r *echo.Group) {
	r.GET("/:id", h.absen)
	r.POST("", h.kirimabsen)
}

func (h *Handler) absen(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	if cc, m := Getabsen(id); m == nil {
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
