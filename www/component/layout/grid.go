package layout

import (
	"github.com/golangee/forms-example/www/component/base"
	"github.com/golangee/forms-example/www/component/view"
	"github.com/golangee/forms-example/www/html"
)

type Grid struct {
	*view.Component
}

func NewGrid() *Grid {
	c := &Grid{}
	c.Component = view.NewComponent("div")
	html.Class("grid grid-cols-3 gap-4")(c.Element())
	return c
}
