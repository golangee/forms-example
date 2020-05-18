// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package table

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
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
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
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
