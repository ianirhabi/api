package engine

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"retrobarbershop.com/retro/api/src/absen"
	"retrobarbershop.com/retro/api/src/barang"
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
	endpoin_absen.Use(middleware.JWT([]byte("secret")))
	endpoin_absen.GET("/:id", absen.Absendetail)
	endpoin_absen.GET("/:id/:from/:to", absen.Getabsendate)
	endpoin_absen.POST("", absen.Kirimabsen)
	endpoin_absen.POST("/izin", absen.Kirimzin)
	// untuk absen
	// user := e.Group("")
	// user.GET("/", usr.Getuser)

	// v := e.Group("/auth")
	// v.POST("/login", auth.Login)

	endpoin_user_login := e.Group("/retrobarbershop/user")
	endpoin_user_login.POST("/login", user.Login)

	endpoin_user := e.Group("/retrobarbershop/user")
	//r.Use(middleware.JWT([]byte("secret")))
	endpoin_user.Use(middleware.JWT([]byte("secret")))
	endpoin_user.GET("", user.GetAllUser)
	endpoin_user.POST("", user.Adduser)
	endpoin_user.DELETE("/:id", user.Del)
	endpoin_user.GET("/:id", user.GetUser)
	endpoin_user.PUT("/:id", user.UpdateUser)

	endpoin_getimage := e.Group("/retrobarbershop/getimage")
	endpoin_getimage.GET("/:imagefile", upload.Getphoto)

	endpoin_upload := e.Group("/retrobarbershop/upload")
	endpoin_upload.Use(middleware.JWT([]byte("secret")))
	endpoin_upload.POST("", upload.CreatePhotoweb)
	endpoin_upload.POST("/android", upload.CreatePhotoandroid)
	endpoin_upload.POST("/namephoto", upload.CreatenamePhoto)

	enpoin_barang := e.Group("/retrobarbershop/barang")
	enpoin_barang.Use(middleware.JWT([]byte("secret")))
	enpoin_barang.GET("/:usergrup/:page", barang.Getbarang)
	enpoin_barang.POST("/:usergrup/:user_id", barang.Postbarang)
	enpoin_barang.PUT("/:usergrup/:user_id/:idbarang", barang.Updatebarang)
	enpoin_barang.DELETE("/:usergrup/delete/:barangid", barang.Deletebarang)

	e.Logger.Fatal(e.Start(":4500"))
}
