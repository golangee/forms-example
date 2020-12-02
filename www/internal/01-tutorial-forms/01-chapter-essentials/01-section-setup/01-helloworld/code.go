package t01helloworld

import (
	"github.com/golangee/forms-example/www/forms/router"
	"github.com/golangee/forms-example/www/forms/view"
)

const Path = "/tutorial/01-helloworld"

func FromQuery(router.Query) view.Renderable {
	return view.Span(view.Text("hello world"))
}
