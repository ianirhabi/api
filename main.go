package main

import (
	"github.com/alfatih/irhabi/irhabi"
	//"github.com/astaxie/beego/orm"

	"github.com/alfatih/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"retrobarbershop.com/retro/api/engine"
	"retrobarbershop.com/retro/api/model"
)

//type (
//	user struct {
//		ID   int    `json:"id"`
//		Name string `json:"name"`
//	}
//)

//var (
//	users = map[int]*user{}
//	seq   = 1
//)

//----------
// Handlers
//----------

//func createUser(c echo.Context) error {
//	u := &user{
//		ID: seq,
//	}
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	users[u.ID] = u
//	seq++
//	return c.JSON(http.StatusCreated, u)
//}

//func getUser(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	return c.JSON(http.StatusOK, users[id])
//}

//func updateUser(c echo.Context) error {
//	u := new(user)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	id, _ := strconv.Atoi(c.Param("id"))
//	users[id].Name = u.Name
//	return c.JSON(http.StatusOK, users[id])
//}

//func deleteUser(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	delete(users, id)
//	return c.NoContent(http.StatusNoContent)
//}
func init() {
	orm.RegisterModel(new(model.User))
	orm.RegisterDataBase("default", "mysql", "root:'ian!@#$%^'@tcp(18.188.151.231:3306)/retrobarbershop_app?charset=utf8")
}

func main() {
	// Start server
	irhabi.StartServer(engine.Router())
}

/**
  tes dengan aplikasi curl  mngcreate user
  curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe Smith"}' \
  localhost:1323/users

  get user
  curl localhost:1323/users/1

  delete user

  curl -X DELETE localhost:1323/users/1
*/
