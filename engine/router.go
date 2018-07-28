package engine

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"retrobarbershop.com/retro/api/src/absen"
	"retrobarbershop.com/retro/api/src/upload"
	"retrobarbershop.com/retro/api/src/user"
)

func Router() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted group
	endpoin_absen := e.Group("/retrobarbershop/absen")
	//r.Use(middleware.JWT([]byte("secret")))
	endpoin_absen.GET("/:id", absen.Absendetail)
	endpoin_absen.GET("/:id/:from/:to", absen.Getabsendate)
	endpoin_absen.POST("", absen.Kirimabsen)
	// untuk absen
	// user := e.Group("")
	// user.GET("/", usr.Getuser)

	// v := e.Group("/auth")
	// v.POST("/login", auth.Login)

	endpoin_user := e.Group("/retrobarbershop/user")
	//r.Use(middleware.JWT([]byte("secret")))
	endpoin_user.POST("/login", user.Login)
	endpoin_user.DELETE("/user/:id", user.Del)
	endpoin_user.GET("/:id", user.GetUser)
	endpoin_user.GET("", user.GetAllUser)
	endpoin_user.PUT("/:id", user.UpdateUser)

	endpoin_upload := e.Group("/retrobarbershop/upload")
	endpoin_upload.POST("", upload.CreatePhotoweb)
	endpoin_upload.POST("/android", upload.CreatePhotoandroid)
	endpoin_upload.GET("/:imagefile", upload.Getphoto)
	endpoin_upload.POST("/namephoto", upload.CreatenamePhoto)

	e.Logger.Fatal(e.Start(":4500"))
}
