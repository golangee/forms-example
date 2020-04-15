package tabview

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
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
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
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
