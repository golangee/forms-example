package textarea

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/textarea"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("textarea").
			Style(Font(Headline1)),
		NewText("Text areas are like text fields but usually only used for "+
			"multiline input.").
			Style(Font(Body), Repel()),

		NewText("Default"),
		NewTextArea().
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default prefilled"),
		NewTextArea().
			SetText("hello\nworld.").
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default with max length"),
		NewTextArea().
			SetLabel("your multiline input").
			SetMaxLength(10).
			Style(Repel()),

		NewText("disabled"),
		NewTextArea().
			SetLabel("your multiline input").
			SetEnabled(false).
			Style(Repel()),

		NewText("invalid"),
		NewTextArea().
			SetLabel("your multiline input").
			SetInvalid(true).
			Style(Repel()),

		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package textarea

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/textarea"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("textarea").
			Style(Font(Headline1)),
		NewText("Text areas are like text fields but usually only used for "+
			"multiline input.").
			Style(Font(Body), Repel()),

		NewText("Default"),
		NewTextArea().
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default prefilled"),
		NewTextArea().
			SetText("hello\nworld.").
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default with max length"),
		NewTextArea().
			SetLabel("your multiline input").
			SetMaxLength(10).
			Style(Repel()),

		NewCode(GoSyntax, code),
	)}
}`
