package download

import "github.com/labstack/echo"

func GetApk(c echo.Context) (e error) {
	return c.File("/home/ubuntu/go/src/retrobarbershop.com/retro/api/download/retrobarbershop-development.apk")
}
