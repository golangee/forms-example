package text

import (
	"github.com/golangee/forms-example/www/forms/router"
	"github.com/golangee/forms-example/www/forms/text"
	"github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/forms/material/icon"
	"github.com/golangee/forms-example/www/forms/tailwindcss/style"
)

const Path = "/tutorial/text"

type ContentView struct {
	*text.Text
}

func NewContentView() *ContentView {
	return &ContentView{
		text.NewText("turtle rock",
			view.Class(
				style.FontBold,
				style.TextGreen600,
				style.P4,
			),
			view.I(view.Class(style.MaterialIcons,style.TextXl),view.Text(icon.AccessAlarm)),

		),
	}
}

func FromQuery(router.Query) view.Renderable {
	return NewContentView()
}