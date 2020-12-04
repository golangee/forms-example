package c02s03s01

import (
	"github.com/golangee/forms-example/www/forms/button"
	"github.com/golangee/forms-example/www/forms/modal"
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

const Path = "/c02s03s01"

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{}
}

func (c *ContentView) Render() Node {
	return Div(Class("w-screen h-screen grid"),
		With(

			button.NewTextButton(
				func() {
					log.NewLogger().Print(ecs.Msg("wut?"))
					modal.ShowDialog(
						Div(
							Text("hello dialog"),
						),
					)
				}).
				SetCaption(Text("show dialog")),

			AddClass("m-auto"),
		),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
