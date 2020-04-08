package drawer

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/drawer"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Drawer").Style(Font(Headline1)),
		NewText("The navigation drawer is a helper to directly point "+
			" to specific states and screens in your app.").Style(Font(Body)),


		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `func (a *App) SetView(v View) {
	a.Application.SetView(
		NewDrawer(
			NewTopAppBar().
				SetTitle("wtk demo").
				SetNavigation(icon.Menu, nil).
				AddActions(NewIconItem(icon.FileDownload, "download", nil)),
			v),
	)
}`
