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
}

func (h *Handler) absen(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("debig disini ", id)
	if cc, m := Getabsen(id); m == nil {
		return c.JSON(http.StatusOK, &cc)
	}
	return e
}
