package engine

import (
	"github.com/alfatih/irhabi/irhabi"
	"github.com/labstack/echo"
)

type RouteHandlers interface {
	URLMapping(r *echo.Group)
}

var handlers = map[string]RouteHandlers{}

func Router() *echo.Echo {
	engine := irhabi.New()
	v := engine.Group("v1/")
	if len(handlers) > 0 {
		for p, h := range handlers {
			h.URLMapping(v.Group(p))
		}
	}
	return engine
}
