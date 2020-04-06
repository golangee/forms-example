package dialog

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/dialog"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("dialogs").Style(Font(Headline1)),
		NewText("Dialogs interact with a user about a task and usually "+
			" require important or destructive decisions.").Style(Font(Body)),
		NewButton("show dialog").
			AddClickListener(func(v View) {
				NewDialog().
					SetTitle("are you sure?").
					SetBody(NewVStack().AddViews(
						NewText("this could go wrong"),
						NewTextField().SetLabel("enter something"),
					)).
					AddAction("perhaps", func(dlg *Dialog) {
						dlg.Close()
					}).
					AddAction("may be", func(dlg *Dialog) {
						dlg.Close()
					}).
					Show(v)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package dialog

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/dialog"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("dialogs").Style(Font(Headline1)),
		NewText("Dialogs interact with a user about a task and usually "+
			" require important or destructive decisions.").Style(Font(Body)),
		NewButton("show dialog").
			AddClickListener(func(v View) {
				NewDialog().
					SetTitle("are you sure?").
					SetBody(NewVStack().AddViews(
						NewText("this could go wrong"),
						NewTextField().SetLabel("enter something"),
					)).
					AddAction("perhaps", func(dlg *Dialog) {
						dlg.Close()
					}).
					AddAction("may be", func(dlg *Dialog) {
						dlg.Close()
					}).
					Show(v)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)}
}`
