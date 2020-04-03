package index

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk-example/demo/button"
	"github.com/worldiety/wtk-example/demo/typography"
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
			"for go and wasm.").Style(Font(Body)),
		NewButton("buttons").AddClickListener(func(v View) {
			v.Context().Navigate(button.Path)
		}).Style(Padding()),
		NewButton("typography").AddClickListener(func(v View) {
			v.Context().Navigate(typography.Path)
		}).Style(Padding()),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}
