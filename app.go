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

package example

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms-example/build"
	"github.com/golangee/forms-example/demo/button"
	"github.com/golangee/forms-example/demo/card"
	"github.com/golangee/forms-example/demo/checkbox"
	"github.com/golangee/forms-example/demo/circularprogress"
	"github.com/golangee/forms-example/demo/dialog"
	"github.com/golangee/forms-example/demo/drawer"
	"github.com/golangee/forms-example/demo/grid"
	"github.com/golangee/forms-example/demo/hstack"
	"github.com/golangee/forms-example/demo/i18n"
	"github.com/golangee/forms-example/demo/index"
	"github.com/golangee/forms-example/demo/linearprogress"
	"github.com/golangee/forms-example/demo/link"
	"github.com/golangee/forms-example/demo/list"
	"github.com/golangee/forms-example/demo/menu"
	"github.com/golangee/forms-example/demo/notfound"
	"github.com/golangee/forms-example/demo/picker"
	"github.com/golangee/forms-example/demo/smallcontent"
	"github.com/golangee/forms-example/demo/snackbar"
	"github.com/golangee/forms-example/demo/table"
	"github.com/golangee/forms-example/demo/tabview"
	"github.com/golangee/forms-example/demo/textarea"
	"github.com/golangee/forms-example/demo/textfield"
	"github.com/golangee/forms-example/demo/theme"
	"github.com/golangee/forms-example/demo/topappbar"
	"github.com/golangee/forms-example/demo/typography"
	"github.com/golangee/forms-example/demo/vstack"
	"github.com/golangee/forms/theme/material/icon"
	"path/filepath"
	"time"
)

type App struct {
	*Application
}

func NewApp() *App {
	a := &App{}
	a.Application = NewApplication(a, build.Env().String())
	return a
}

func (a *App) WithDrawer(f func(q Query) View) func(Query) View {
	return func(query Query) View {
		v := NewGroup(f(query)).Style(Padding())

		var items []LstItem
		items = append(items,
			NewListItem("home").SetLeadingView(NewIcon(icon.Home)).AddClickListener(func(v View) {
				a.Context().Navigate("/")
			}),
			NewListSeparator(),
			NewListHeader("components"),
		)
		for _, route := range a.Context().Routes() {
			fPath := route.Path
			name := filepath.Base(route.Path)
			if fPath == "/" {
				continue
			}
			item := NewListItem(name)
			if route.Path == query.Path() {
				item.SetSelected(true)
			}
			items = append(items, item.AddClickListener(func(v View) {
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
				AddActions(NewIconItem(icon.Help, "download", func(v View) {
					ShowMessage(v, "wtk demo")
				})),
			NewVStack().AddViews(
				NewText("your demo").Style(Font(DrawerTitle)),
				NewText("anonymous").Style(Font(DrawerSubTitle)),
			),
			NewList().AddItems(items...),
			v)
	}
}

func (a *App) Start() {
	a.UnmatchedRoute(notfound.FromQuery)
	a.Route(index.Path, a.WithDrawer(index.FromQuery))
	a.Route(button.Path, a.WithDrawer(button.FromQuery))
	a.Route(typography.Path, a.WithDrawer(typography.FromQuery))
	a.Route(textfield.Path, a.WithDrawer(textfield.FromQuery))
	a.Route(textarea.Path, a.WithDrawer(textarea.FromQuery))
	a.Route(dialog.Path, a.WithDrawer(dialog.FromQuery))
	a.Route(menu.Path, a.WithDrawer(menu.FromQuery))
	a.Route(topappbar.Path, a.WithDrawer(topappbar.FromQuery))
	a.Route(drawer.Path, a.WithDrawer(drawer.FromQuery))
	a.Route(list.Path, a.WithDrawer(list.FromQuery))
	a.Route(link.Path, a.WithDrawer(link.FromQuery))
	a.Route(theme.Path, a.WithDrawer(theme.FromQuery))
	a.Route(table.Path, a.WithDrawer(table.FromQuery))
	a.Route(checkbox.Path, a.WithDrawer(checkbox.FromQuery))
	a.Route(picker.Path, a.WithDrawer(picker.FromQuery))
	a.Route(tabview.Path, a.WithDrawer(tabview.FromQuery))
	a.Route(linearprogress.Path, a.WithDrawer(linearprogress.FromQuery))
	a.Route(circularprogress.Path, a.WithDrawer(circularprogress.FromQuery))
	a.Route(snackbar.Path, a.WithDrawer(snackbar.FromQuery))
	a.Route(smallcontent.Path, a.WithDrawer(smallcontent.FromQuery))
	a.Route(card.Path, a.WithDrawer(card.FromQuery))
	a.Route(i18n.Path, a.WithDrawer(i18n.FromQuery))
	a.Route(grid.Path, a.WithDrawer(grid.FromQuery))
	a.Route(vstack.Path, a.WithDrawer(vstack.FromQuery))
	a.Route(hstack.Path, a.WithDrawer(hstack.FromQuery))
	a.Application.Start()
}
