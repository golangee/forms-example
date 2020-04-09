package index

import (
	. "github.com/worldiety/wtk"
)

const Path = "/"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("wtk").Style(Font(Headline1)).Style(PadBottom(8)),
		NewText("welcome to the wtk demo and kitchen sink. Here you can see "+
			"a selected amount of use cases and components of the worldiety web toolkit "+
			"for go and wasm. It should feel a bit like SwiftUI but for Go.").Style(Font(Body)),
		NewText("goals").Style(Font(Headline2)),
		NewText("Provide an 80% solution for rapid prototyping and "+
			"digitization of companies. The UI must be state of the art "+
			"and working across desktop and mobile, in modern Webkit and Firefox "+
			"Browsers. Best suited for form based web apps.").Style(Font(Body)),
		NewText("non-goals").Style(Font(Headline2)),
		NewText("Developing bleeding edge and fully customizable html "+
			"applications. Compatibility with older browsers or SEO "+
			"are unimportant. Everything should be expressed in a declarative way "+
			"without the need of external Javascript libraries.").Style(Font(Body)),

	)}
}

func FromQuery(Query) View {
	return NewContentView()
}
