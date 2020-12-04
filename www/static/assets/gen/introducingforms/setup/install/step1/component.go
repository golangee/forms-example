package c01s01s01

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
)

const Path = "/c01s01s01"

func Show(q router.Query) Renderable {
	return Div(Text("Hello world4"))
}
