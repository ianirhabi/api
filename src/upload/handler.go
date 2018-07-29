package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"

	"github.com/labstack/echo"
)

// type Handler struct{}

// func (h *Handler) URLMapping(r *echo.Group) {
// 	r.POST("", h.createPhotoweb)
// 	r.POST("/android", h.createPhotoandroid)
// 	r.GET("/:imagefile", h.getphoto)
// 	r.POST("/namephoto", h.createnamePhoto)
// }

// untuk web
func CreatePhotoweb(c echo.Context) (e error) {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("upload/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

type lod struct {
	Upload string `json:"upload"`
	ID     int64  `json:"id"`
	Resp   string `json:"respon,omitempty"`
}

var b lod

//untuk android
func CreatePhotoandroid(c echo.Context) (e error) {
	// Read form fields

	name := c.FormValue("name")
	email := c.FormValue("email")

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["cycle"]

	diks := form.File["diran"]
	dd := diks[0]
	a := dd.Filename
	bb := dd.Header

	fmt.Println("disknya", a, " ", bb)
	fmt.Println("nama ", b.Upload)

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("/root/go/src/retrobarbershop.com/retro/api/upload/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))
}

func CreatenamePhoto(c echo.Context) (e error) {
	if err := c.Bind(&b); err == nil {
		var uon model.User
		o := orm.NewOrm()

		if d := o.Raw("update user set nama_foto = ? where id = ?", b.Upload, b.ID).QueryRow(&uon); d != nil {
			b.Resp = "Berhasil Upload"
			return c.JSON(http.StatusCreated, &b)
		} else {
			b.Resp = "Maaf Anda Gagal Upload"
			return c.JSON(http.StatusCreated, &b)
		}
	}
	return e
}

func Getphoto(c echo.Context) (e error) {
	ifile, _ := strconv.Atoi(c.Param("imagefile"))
	var a = strconv.Itoa(ifile)
	return c.File("/root/go/src/retrobarbershop.com/retro/api/upload/" + a + ".jpg")
}
