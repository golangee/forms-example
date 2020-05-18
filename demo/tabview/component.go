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

package tabview

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
)

const Path = "/demo/tabview"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Tabview").Style(Font(Headline1)),
		NewText("Tabs provide navigational access to groups of content "+
			"which is related to each other at the same level of hierarchy.").Style(Font(Body)),

		NewTabView().SetTabs(
			NewTab("tab 1", NewText("content 1")),
			NewTab("tab 2", NewText("content 2")),
		).Style(Repel()),

		NewTabView().SetTabs(
			NewTabWithIcon(icon.Favorite, "tab 1", NewText("content 1")),
			NewTabWithIcon(icon.Help, "tab 2", NewText("content 2")),
		).SetActive(1).Style(Repel()),

		NewTabView().SetTabs(
			NewTabWithStackedIcon(icon.Favorite, "tab 1", NewText("content 1")),
			NewTabWithStackedIcon(icon.Help, "tab 2", NewText("content 2")),
		).SetActive(1).Style(Repel()),

		NewTabView().SetTabs(
			NewTab("tabulator 1", NewText("content 1")),
			NewTab("tabulator 2", NewText("content 2")),
			NewTab("tabulator 3", NewText("content 3")),
			NewTab("tabulator 4", NewText("content 4")),
			NewTab("tabulator 5", NewText("content 5")),
			NewTab("tabulator 6", NewText("content 6")),
			NewTab("tabulator 7", NewText("content 7")),
		).SetScrollable(true).
			Style(Width(Pixel(400)), Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package tabview

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
)

const Path = "/demo/tabview"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Tabview").Style(Font(Headline1)),
		NewText("Tabs provide navigational access to groups of content "+
			"which is related to each other at the same level of hierarchy.").Style(Font(Body)),

		NewTabView().SetTabs(
			NewTab("tab 1", NewText("content 1")),
			NewTab("tab 2", NewText("content 2")),
		).Style(Repel()),

		NewTabView().SetTabs(
			NewTabWithIcon(icon.Favorite, "tab 1", NewText("content 1")),
			NewTabWithIcon(icon.Help, "tab 2", NewText("content 2")),
		).SetActive(1).Style(Repel()),

		NewTabView().SetTabs(
			NewTabWithStackedIcon(icon.Favorite, "tab 1", NewText("content 1")),
			NewTabWithStackedIcon(icon.Help, "tab 2", NewText("content 2")),
		).SetActive(1).Style(Repel()),

		NewTabView().SetTabs(
			NewTab("tabulator 1", NewText("content 1")),
			NewTab("tabulator 2", NewText("content 2")),
			NewTab("tabulator 3", NewText("content 3")),
			NewTab("tabulator 4", NewText("content 4")),
			NewTab("tabulator 5", NewText("content 5")),
			NewTab("tabulator 6", NewText("content 6")),
			NewTab("tabulator 7", NewText("content 7")),
		).SetScrollable(true).
			Style(Width(Pixel(400)), Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}`
