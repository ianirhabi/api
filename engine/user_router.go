package engine

import "retrobarbershop.com/retro/api/src/user"

func init() {
	handlers["user"] = &user.Handler{}
	//	handlers["user"] = &user.Handler{}
}
