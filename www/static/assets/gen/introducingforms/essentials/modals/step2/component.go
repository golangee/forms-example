package c02s03s02

import (
	"github.com/golangee/forms-example/www/forms/button"
	"github.com/golangee/forms-example/www/forms/modal"
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
)

const Path = "/c02s03s02"

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{}
}

func (c *ContentView) Render() Node {
	return Div(Class("w-screen h-screen grid"),
		modal.NewPopupMenu(
			With(
				button.NewTextButton("show menu", nil),
				AddClass("m-auto"),
			),
			Div(
				modal.NewMenuItem(Text("hey")),
				modal.NewMenuItem(Text("ho")),
			),
		),

	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
