package snackbar

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/snackbar"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Snackbar").Style(Font(Headline1)),
		NewText("A snackbar provides short summarized message and an action option.").Style(Font(Body)),
		NewButton("snack it").AddClickListener(func(v View) {
			NewSnackbar("Here comes a snack.", "Get it").
				SetAction(func(v View) {
					ShowMessage(v, "you got the snack")
				}).
				Show(v)
		}),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package snackbar

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/snackbar"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Snackbar").Style(Font(Headline1)),
		NewText("A snackbar provides short summarized message and an action option.").Style(Font(Body)),
		NewButton("snack it").AddClickListener(func(v View) {
			NewSnackbar("Here comes a snack.", "Get it").
				SetAction(func(v View) {
					ShowMessage(v, "you got the snack")
				}).
				Show(v)
		}),

		NewCode(GoSyntax, code),
	)
	return view
}`
