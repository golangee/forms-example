package theme

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/theme"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	colorButtons := NewGroup()
	for name, color := range Colors {
		myColor := color
		colorButtons.AddViews(
			NewButton(name).
				AddClickListener(func(v View) {
					Theme().SetColor(myColor)
				}).Style(ForegroundColor(myColor)),
		)
	}

	view.VStack = NewVStack().AddViews(
		NewText("Theme").Style(Font(Headline1)),
		NewText("Colors are important to increase the recognition factor "+
			" and to focus on sensitive actions.").Style(Font(Body)),
		colorButtons,
		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package theme

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/theme"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	colorButtons := NewGroup()
	for name, color := range Colors {
		myColor := color
		colorButtons.AddViews(
			NewButton(name).
				AddClickListener(func(v View) {
					Theme().SetPrimaryColor(myColor)
				}).Style(ForegroundColor(myColor)),
		)
	}

	view.VStack = NewVStack().AddViews(
		NewText("Theme").Style(Font(Headline1)),
		NewText("Colors are important to increase the recognition factor "+
			" and to focus on sensitive actions.").Style(Font(Body)),
		colorButtons,
		NewCode(GoSyntax, code),
	)
	return view
}`
