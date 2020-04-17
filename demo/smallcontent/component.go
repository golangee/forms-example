package smallcontent

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/smallcontent"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("small content").Style(Font(Headline1)),
		NewText("Just to test longer menu with small content body and a bit of breaking text.").Style(Font(Body)),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
