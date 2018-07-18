package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alfatih/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"retrobarbershop.com/retro/api/model"
)

type Handler struct{}

type user struct {
	Name  string `json:"username"`
	Passw string `json:"password"`
	Res   string `json:"Res,omitempty"`
	Usr   model.User
}

var (
	users = map[int]*user{}
	seq   = 1
)
var b user

func (h *Handler) URLMapping(r *echo.Group) {
	r.POST("", h.login)
	r.DELETE("/:id", h.del)
	r.GET("/:id", h.getUser)
	r.GET("/userfoto/", h.getUserFoto)
}

func (h *Handler) login(c echo.Context) (e error) {
	b.Res = "tidak boleh"
	if err := c.Bind(&b); err == nil {
		if b.Name == "123" {
			b.Res = "Email 123"
			return c.JSON(http.StatusNonAuthoritativeInfo, &b)
		} else if b.Name == "222" {
			b.Res = "Email 222"
			return c.JSON(http.StatusNonAuthoritativeInfo, &b)
		} else {
			o := orm.NewOrm()

			//user := model.User{User: b.Name,
			//	Pass: b.Passw}
			// insert
			//data, err := o.Insert(&user)
			//fmt.Printf("ID: %d, ERR: %v\n", data, err)
			var uon []*model.User
			o.Raw("SELECT * FROM user").QueryRows(&uon)
			for _, row := range uon {
				b.Usr.Id = row.Id
				b.Usr.User = row.User
				b.Usr.Pass = row.Pass
				b.Usr.Usergup = row.Usergup
				if row.User == b.Name && row.Pass == b.Passw {
					fmt.Println("Berhasil Login")
					b.Res = "ok"
					break
				}
			}
			return c.JSON(http.StatusCreated, &b)
		}
	}
	return e
}

func (h *Handler) del(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) getUser(c echo.Context) (e error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var uon model.User
	if cc, m := Getuser(id, uon); m == nil {
		return c.JSON(http.StatusOK, &cc)
	}
	return e
}

func (h *Handler) getUserFoto(c echo.Context) (e error) {
	return c.File("userfoto/cover.jpg")
}
