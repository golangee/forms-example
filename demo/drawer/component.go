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

package drawer

import (
	. "github.com/golangee/forms"
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

const code = `func (a *App) WithDrawer(f func(q Query) View) func(Query) View {
	return func(query Query) View {
		v := f(query)

		var items []LstItem
		items = append(items, NewListSeparator(), NewListHeader("components"))
		for _, route := range a.Context().Routes() {
			fPath := route.Path
			name := filepath.Base(route.Path)
			if fPath == "/" {
				name = "home"
			}
			items = append(items, NewListItem(name).AddClickListener(func(v View) {
				go func() {
					time.Sleep(200 * time.Millisecond) // wait for drawer animation
					a.Context().Navigate(fPath)
				}()
			}))
		}

		return NewDrawer(
			NewTopAppBar().
				SetTitle("wtk demo").
				SetNavigation(icon.Menu, nil).
				AddActions(NewIconItem(icon.FileDownload, "download", nil)),
			NewVStack().AddViews(
				NewText("your demo").Style(Font(DrawerTitle)),
				NewText("anonymous").Style(Font(DrawerSubTitle)),
			),
			NewList().AddItems(items...),
			v)
	}
}`
