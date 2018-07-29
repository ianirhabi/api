package user

import (
	"time"

	"retrobarbershop.com/retro/api/model"
)

type Requestlogin struct {
	Name  string `json:"username"`
	Passw string `json:"password"`
	Usr   model.User
}
type ResponsAlluser struct {
	Status string      `json:"status,omitempty "`
	Data   interface{} `json:"data"`
}

type Respons struct {
	Status string      `json:"status,omitempty"`
	Jam    time.Time   `json:"Last_login,omitempty"`
	Data   interface{} `json:"data"`
}

type UpdateRequest struct {
	Password string `json:"password"`
}

type AddRequest struct {
	Name       string `json:"nama"`
	User       string `json:"username"`
	Pass       string `json:"password"`
	Usergrup   string `json:"usergrup"`
	Ianmonitor string `json:"controlian"`
}
