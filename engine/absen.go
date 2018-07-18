package engine

import (
	"retrobarbershop.com/retro/api/src/absen"
)

func init() {
	handlers["absen"] = &absen.Handler{}
}
