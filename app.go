package example

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk-example/demo/button"
	"github.com/worldiety/wtk-example/demo/checkbox"
	"github.com/worldiety/wtk-example/demo/circularprogress"
	"github.com/worldiety/wtk-example/demo/dialog"
	"github.com/worldiety/wtk-example/demo/drawer"
	"github.com/worldiety/wtk-example/demo/index"
	"github.com/worldiety/wtk-example/demo/linearprogress"
	"github.com/worldiety/wtk-example/demo/link"
	"github.com/worldiety/wtk-example/demo/list"
	"github.com/worldiety/wtk-example/demo/menu"
	"github.com/worldiety/wtk-example/demo/notfound"
	"github.com/worldiety/wtk-example/demo/picker"
	"github.com/worldiety/wtk-example/demo/snackbar"
	"github.com/worldiety/wtk-example/demo/table"
	"github.com/worldiety/wtk-example/demo/tabview"
	"github.com/worldiety/wtk-example/demo/textarea"
	"github.com/worldiety/wtk-example/demo/textfield"
	"github.com/worldiety/wtk-example/demo/theme"
	"github.com/worldiety/wtk-example/demo/topappbar"
	"github.com/worldiety/wtk-example/demo/typography"
	"github.com/worldiety/wtk/theme/material/icon"
	"path/filepath"
	"time"
)

type App struct {
	*Application
}

func NewApp() *App {
	a := &App{}
	a.Application = NewApplication(a)
	return a
}

func (a *App) WithDrawer(f func(q Query) View) func(Query) View {
	return func(query Query) View {
		v := f(query)

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
	a.Application.Start()
}
