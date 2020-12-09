package c02s03s02

import (
	"github.com/golangee/forms-example/www/forms/button"
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/forms-example/www/forms/menu"
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
	btnId := dom.GenerateID()

	return Div(Class("w-screen h-screen grid"),
		With(button.NewTextButton("show menu", func() {
			menu.ShowPopup(btnId, menu.NewMenu(Div(
				menu.NewMenuItem(Text("hey2")),
				menu.NewMenuItem(Text("ho2")),
			)))
		}), AddClass("m-auto"), ID(btnId)),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
