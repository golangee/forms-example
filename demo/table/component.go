package table

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
	"strconv"
)

const Path = "/demo/table"

type ContentView struct {
	*VStack
	table   *Table
	btn     *Button
	indices *VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Table").Style(Font(Headline1)),
		NewText("A data table presents information in rows and columns. "+
			"Users need to understand and look for patterns and insights.").Style(Font(Body), Repel()),
		NewText("selected rows:"),
		NewVStack().Self(&view.indices).Style(Repel()),
		NewTable().Self(&view.table).
			Align(0, Leading).Align(1, Trailing).Align(2, Trailing).
			SetHeader(NewText("name"), NewText("value"), NewText("price")).
			AddRow(NewText("Bibo"), NewText("123.45"), NewText("12,49 EUR")).
			AddRow(NewText("Samson"), NewText("45.67"), NewText("4,56 EUR")).
			AddRow(NewButton("hello"), NewIcon(icon.Favorite), NewTextField()).
			SetSelectionChangeListener(func(t *Table) {
				view.indices.RemoveAll()
				for _, val := range t.Selected() {
					view.indices.AddViews(NewText(strconv.Itoa(val)))
				}
			}).
			SetSelected(1).
			Style(Repel()),

		NewButton("toggle selection").AddClickListener(func(v View) {
			view.table.SetRowSelection(!view.table.RowSelection())
		}),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package table

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
	"strconv"
)

const Path = "/demo/table"

type ContentView struct {
	*VStack
	table   *Table
	btn     *Button
	indices *VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Table").Style(Font(Headline1)),
		NewText("A data table presents information in rows and columns. "+
			"Users need to understand and look for patterns and insights.").Style(Font(Body), Repel()),
		NewText("selected rows:"),
		NewVStack().Self(&view.indices).Style(Repel()),
		NewTable().Self(&view.table).
			Align(0, Leading).Align(1, Trailing).Align(2, Trailing).
			SetHeader(NewText("name"), NewText("value"), NewText("price")).
			AddRow(NewText("Bibo"), NewText("123.45"), NewText("12,49 EUR")).
			AddRow(NewText("Samson"), NewText("45.67"), NewText("4,56 EUR")).
			AddRow(NewButton("hello"), NewIcon(icon.Favorite), NewTextField()).
			SetSelectionChangeListener(func(t *Table) {
				view.indices.RemoveAll()
				for _, val := range t.Selected() {
					view.indices.AddViews(NewText(strconv.Itoa(val)))
				}
			}).
			SetSelected(1).
			Style(Repel()),

		NewButton("toggle selection").AddClickListener(func(v View) {
			view.table.SetRowSelection(!view.table.RowSelection())
		}),

		NewCode(GoSyntax, code),
	)
	return view
}`
