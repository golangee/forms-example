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

package list

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"strconv"
)

const Path = "/demo/list"

type ContentView struct {
	*VStack
	btn           *Button
	selectMsg     *Text
	selectionList *List
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("List").Style(Font(Headline1)),
		NewText("Lists are vertical elements of text or images. ").Style(Font(Body)),

		NewList().AddItems(
			NewListItem("line 0").
				AddClickListener(func(v View) {
					ShowMessage(v, "clicked line 0")
				}),
			NewListItem("line 1"),
			NewListItem("line 2"),
			NewListSeparator(),
			NewListHeader("now with icons"),
			NewListItem("line 3").SetLeadingView(NewIcon(icon.Add)),
			NewListItem("line 3").SetLeadingView(NewIcon(icon.Add)),
		).Style(Repel()),

		NewText("nothing selected").Self(&view.selectMsg),
		NewButton("select Option 2").AddClickListener(func(v View) {
			view.selectionList.SetSelectedIndex(1)
		}).Style(Repel()),

		NewSelectionList().Self(&view.selectionList).AddItems(
			NewListItem("Option 1"),
			NewListItem("Option 2"),
			NewListItem("Option 3"),
		).AddSelectListener(func(idx int) {
			view.selectMsg.Set("selected option " + strconv.Itoa(idx))
		}).Style(Repel()),

		NewList().AddItems(
			NewListItem("cool").
				SetLeadingView(NewIcon(icon.Add)).
				SetTrailingView(
					NewButton("").
						AddIcon(icon.MoreVert, Leading).
						AddClickListener(func(v View) {
							ShowMenu(v, NewMenuItem("awesome", nil))
						}),
				),
		).Style(Repel()),

		NewList().AddItems(
			NewListTwoLineItem("important information", "with some more details"),
			NewListTwoLineItem("hello", "world"),
			NewListTwoLineItem("hello", "world").SetLeadingView(NewIcon(icon.Favorite).Style(ForegroundColor(RGBA(255, 0, 0, 255)))),
			NewListTwoLineItem("hello", "world").SetLeadingView(NewIcon(icon.Add)).SetTrailingView(NewIcon(icon.SkipNext)),
		).Style(Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package list

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"strconv"
)

const Path = "/demo/list"

type ContentView struct {
	*VStack
	btn           *Button
	selectMsg     *Text
	selectionList *List
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("List").Style(Font(Headline1)),
		NewText("Lists are vertical elements of text or images. ").Style(Font(Body)),

		NewList().AddItems(
			NewListItem("line 0").
				AddClickListener(func(v View) {
					ShowMessage(v, "clicked line 0")
				}),
			NewListItem("line 1"),
			NewListItem("line 2"),
			NewListSeparator(),
			NewListHeader("now with icons"),
			NewListItem("line 3").SetIcon(icon.Add),
			NewListItem("line 3").SetIcon(icon.Favorite).SetSelected(true),
		).Style(Repel()),

		NewText("nothing selected").Self(&view.selectMsg),
		NewButton("select Option 2").AddClickListener(func(v View) {
			view.selectionList.SetSelectedIndex(1)
		}).Style(Repel()),

		NewSelectionList().Self(&view.selectionList).AddItems(
			NewListItem("Option 1"),
			NewListItem("Option 2"),
			NewListItem("Option 3"),
		).AddSelectListener(func(idx int) {
			view.selectMsg.Set("selected option " + strconv.Itoa(idx))
		}).Style(Repel()),

		NewList().AddItems(
			NewListItem("cool").
				SetIcon(icon.Call).
				SetTrailingView(
					NewButton("").
						AddIcon(icon.MoreVert, Leading).
						AddClickListener(func(v View) {
							ShowMenu(v, NewMenuItem("awesome", nil))
						}),
				),

		),

		NewCode(GoSyntax, code),
	)
	return view
}`
