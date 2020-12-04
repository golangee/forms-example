package c01s01s01

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
)

func Show(q router.Query) Renderable {
	return Div(Text("Hello world4"))
}
