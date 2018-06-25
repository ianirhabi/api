package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	var d c
	d.a = "kamu"
	d.b = "dia"
	d.cc = "bersamanya"
	var dis = []string{d.a, " ", d.b, " ", d.cc, " "}

	e.GET("/", func(c echo.Context) error {
		for _, ff := range dis {
			c.String(http.StatusOK, ff)
		}
		return c.String(http.StatusOK, "mulai")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type c struct {
	a  string
	b  string
	cc string
}
