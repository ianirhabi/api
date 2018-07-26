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
