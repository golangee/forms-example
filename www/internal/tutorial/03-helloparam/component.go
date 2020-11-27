// Package t03helloparam shows how to create a custom hello world component, instead of just using a function as
// shown in package t01helloworld. To do so, just aggregate a view.View into your custom struct and implement
// a 'func Render() view.Node' method, so that your new component actually conforms to the view.Component interface.
// Note also the usage of the query parameter 'name' to say hello.
package t03helloparam

import (
	"github.com/golangee/forms-example/www/forms/router"
	"github.com/golangee/forms-example/www/forms/view"
)

const Path = "/tutorial/03-helloworld?name=world"

type ContentView struct {
	msg string
	view.View
}

func NewContentView(msg string) *ContentView {
	return &ContentView{
		msg: msg,
	}
}

func (c *ContentView) Render() view.Node {
	return view.Span(view.Text("hello " + c.msg))
}

func FromQuery(q router.Query) view.Renderable {
	return NewContentView(q.Str("name"))
}
