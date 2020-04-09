package link

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/link"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Link").Style(Font(Headline1)),
		NewText("A simple inline link "+
			" for text based non-button and non-icon navigation.").Style(Font(Body)),
		NewGroup(
			NewText("hello "),
			NewLink("world", "http://www.worldiety.de").SetTarget(TargetBlank),
		),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package link

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/link"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Link").Style(Font(Headline1)),
		NewText("A simple inline link "+
			" for text based non-button and non-icon navigation.").Style(Font(Body)),
		NewGroup(
			NewText("hello "),
			NewLink("world", "http://www.worldiety.de").SetTarget(TargetBlank),
		),

		NewCode(GoSyntax, code),
	)
	return view
}`
