package index

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk-example/demo/button"
	"github.com/worldiety/wtk-example/demo/dialog"
	"github.com/worldiety/wtk-example/demo/menu"
	"github.com/worldiety/wtk-example/demo/textarea"
	"github.com/worldiety/wtk-example/demo/textfield"
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
		NewText("chapters").Style(Font(Headline2)),
		NewButton("buttons").AddClickListener(func(v View) {
			v.Context().Navigate(button.Path)
		}).Style(Padding()),
		NewButton("typography").AddClickListener(func(v View) {
			v.Context().Navigate(typography.Path)
		}).Style(Padding()),
		NewButton("textfields").AddClickListener(func(v View) {
			v.Context().Navigate(textfield.Path)
		}).Style(Padding()),
		NewButton("textareas").AddClickListener(func(v View) {
			v.Context().Navigate(textarea.Path)
		}).Style(Padding()),
		NewButton("dialogs").AddClickListener(func(v View) {
			v.Context().Navigate(dialog.Path)
		}).Style(Padding()),
		NewButton("menu").AddClickListener(func(v View) {
			v.Context().Navigate(menu.Path)
		}).Style(Padding()),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}
