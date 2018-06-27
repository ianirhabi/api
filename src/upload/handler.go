package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alfatih/beego/orm"
	"retrobarbershop.com/retro/api/model"

	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) URLMapping(r *echo.Group) {
	r.POST("", h.createPhotoweb)
	r.POST("/android", h.createPhotoandroid)
	r.GET("/", h.getphoto)
	r.POST("/namephoto", h.createnamePhoto)
}

func (h *Handler) createPhotoweb(c echo.Context) (e error) {
	fmt.Println("Masuk Pertama")
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

	fmt.Println("filenya", file)

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
	Resp   string
	ID     int64 `json:"id"`
}

var b lod

func (h *Handler) createPhotoandroid(c echo.Context) (e error) {
	// Read form fields
	fmt.Println("Masuk android")
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

	fmt.Println("filenya", files[0])

	fmt.Println("nama ", b.Upload)

	for _, file := range files {
		// Source
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

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))
}

func (h *Handler) createnamePhoto(c echo.Context) (e error) {
	if err := c.Bind(&b); err == nil {
		fmt.Println("masuk id nama foto "+" ini id nya ", b.ID, " ini nama foto ", b.Upload)
		uon := new(model.User)
		o := orm.NewOrm()
		o.Raw("update user set nama_foto = ? where id = ?", b.Upload, b.ID).QueryRow(uon)
		return c.JSON(http.StatusCreated, &b)
	}
	return e
}

func (h *Handler) getphoto(c echo.Context) (e error) {
	fmt.Println("welcome")
	return c.File("upload/1530071854757.jpg")
}
