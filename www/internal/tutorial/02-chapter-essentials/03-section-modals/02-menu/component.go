package c02s03s02

import (
	"github.com/golangee/forms-example/www/forms/button"
	"github.com/golangee/forms-example/www/forms/layout"
	"github.com/golangee/forms-example/www/forms/modal"
	"github.com/golangee/forms-example/www/forms/popper"
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
	var p *popper.Popper
	return Div(Class("w-screen h-screen grid"),
		/*
			With(
				modal.NewPopupMenu(

					button.NewTextButton("show menu", nil),


					Div(
						modal.NewMenuItem(Text("hey")),
						modal.NewMenuItem(Text("ho")),
					),
				),
				AddClass("m-auto"),
			)*/
		popper.NewPopper(button.NewTextButton("show menu",
			func() {
				p.ShowProperty().Toggle()
			},
		),
			layout.NewVStack(
				modal.NewMenuItem(Text("hey")),
				modal.NewMenuItem(Text("ho")),
			),
		).Self(&p),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
