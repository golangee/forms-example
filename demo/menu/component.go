package menu

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/menu"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("menu").Style(Font(Headline1)),
		NewText("Menu shows a list of actions in a dialog-like "+
			" container.").Style(Font(Body)),
		NewButton("show menu").Self(&view.btn).
			AddClickListener(func(v View) {
				ShowMenu(v,
					NewMenuItem("first item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the first")
					}),
					NewMenuItem("second item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the second")
					}),
					NewMenuDivider(),
					NewMenuItem("third item", nil).SetEnabled(false),
				)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package menu

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/menu"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("menu").Style(Font(Headline1)),
		NewText("Menu shows a list of actions in a dialog-like "+
			" container.").Style(Font(Body)),
		NewButton("show menu").Self(&view.btn).
			AddClickListener(func(v View) {
				ShowMenu(v,
					NewMenuItem("first item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the first")
					}),
					NewMenuItem("second item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the second")
					}),
					NewMenuDivider(),
					NewMenuItem("third item", nil).SetEnabled(false),
				)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)
	return view
}`
