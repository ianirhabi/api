package engine

import (
	"retrobarbershop.com/retro/api/src/upload"
)

func init() {
	handlers["upload"] = &upload.Handler{}
}
