package example

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk-example/demo/button"
	"github.com/worldiety/wtk-example/demo/dialog"
	"github.com/worldiety/wtk-example/demo/drawer"
	"github.com/worldiety/wtk-example/demo/index"
	"github.com/worldiety/wtk-example/demo/list"
	"github.com/worldiety/wtk-example/demo/menu"
	"github.com/worldiety/wtk-example/demo/notfound"
	"github.com/worldiety/wtk-example/demo/textarea"
	"github.com/worldiety/wtk-example/demo/textfield"
	"github.com/worldiety/wtk-example/demo/topappbar"
	"github.com/worldiety/wtk-example/demo/typography"
	"github.com/worldiety/wtk/theme/material/icon"
	"path/filepath"
	"strconv"
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

type MyCustomComponent struct {
	*VStack
	text1   *Text
	text2   *Text
	counter int
}

func NewMyCustomComponent() *MyCustomComponent {
	c := &MyCustomComponent{}

	NewVStack().AddViews(
		NewText("hello world 2").Self(&c.text1).Style(PadLeft(12), ForegroundColor(RGBA(255, 0, 0, 255))),
		NewText("a second text line").Self(&c.text2),
		NewButton("press me").AddClickListener(func(v View) {
			c.counter++
			c.text1.Set("text 1: clicked " + strconv.Itoa(c.counter))
			c.text2.Set("text 2: clicked " + strconv.Itoa(c.counter))

			text := NewText("your click no " + strconv.Itoa(c.counter))
			c.VStack.AddViews(text)
		}),
		NewButton("button 2").AddClickListener(func(v View) {
			v.(*Button).SetEnabled(false)
		}).SetStyleKind(Raised),
		NewButton("button 3").SetStyleKind(Outlined).AddIcon(icon.AccessAlarm, Trailing),
		NewButton("button 4").SetStyleKind(Unelevated).AddIcon(icon.AddAlert, Leading),
		NewButton("button 5").SetStyleKind(Default).AddIcon(icon.Call, Leading).AddIcon(icon.Close, Trailing),
	).Self(&c.VStack)

	return c
}

func (a *App) WithDrawer(f func(q Query) View) func(Query) View {
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
				AddActions(NewIconItem(icon.FileDownload, "download", nil)),
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
	a.Application.Start()
}
