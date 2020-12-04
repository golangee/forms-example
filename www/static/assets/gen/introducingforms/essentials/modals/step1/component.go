package c02s03s01

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
)

const Path = "/c02s03s01"

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{}
}

func (c *ContentView) Render() Node {
	return Div(Text("hello dialog 2"))
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
