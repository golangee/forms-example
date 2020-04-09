package table

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/table"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Table").Style(Font(Headline1)),
		NewText("A data table presents information in rows and columns. "+
			"Users need to understand and look for patterns and insights.").Style(Font(Body)),
		NewTable().
			Align(0, Leading).Align(1, Trailing).Align(2, Trailing).
			SetHeader(NewText("name"), NewText("value"), NewText("price")).
			AddRow(NewText("Bibo"), NewText("123.45"), NewText("12,49 EUR")).
			AddRow(NewText("Samson"), NewText("45.67"), NewText("4,56 EUR")).Layout(),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = ``
