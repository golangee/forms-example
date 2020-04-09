package topappbar

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/topappbar"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Top App Bar").Style(Font(Headline1)),
		NewText("The top app bar is a container for navigation (menu, back or up), "+
			" application title and actions.").Style(Font(Body)),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package topappbar

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
	"log"
)

const Path = "/demo/topappbar"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Top App Bar").Style(Font(Headline1)),
		NewText("The top app bar is a container for navigation (menu, back or up), "+
			" application title and actions.").Style(Font(Body)),
		NewTopAppBar().
			SetNavigation(icon.Menu, func(view View) {
				log.Println("pressed the menu")
			}).
			SetTitle("my action bar").
			AddActions(
				NewIconItem(icon.FileDownload, "download", func(v View) {
					log.Println("download")
				}),
				NewIconItem(icon.Print, "print", func(v View) {
					log.Print("print")
				}),
			),

		NewCode(GoSyntax, code),
	)
	return view
}`
