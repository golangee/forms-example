// Package t01helloworld shows the obligatory 'hello world' example, by performing just do some client side html
// rendering, there is no need to create a component at all.
package t01helloworld

import (
	"github.com/golangee/forms-example/www/forms/router"
	"github.com/golangee/forms-example/www/forms/view"
)

const Path = "/tutorial/01-helloworld"

func FromQuery(router.Query) view.Renderable {
	return view.Span(view.Text("hello world"))
}
