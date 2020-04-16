package circularprogress

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/circularprogress"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Circular Progress").Style(Font(Headline1)),
		NewText("An entertaining and progress indicator component in a circle.").Style(Font(Body)),

		NewCircularProgress(),


		NewCode(GoSyntax, code),
	)

	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package circularprogress

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/circularprogress"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Circular Progress").Style(Font(Headline1)),
		NewText("An entertaining and progress indicator component in a circle.").Style(Font(Body)),

		NewCircularProgress(),


		NewCode(GoSyntax, code),
	)

	return view
}`
