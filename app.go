package example

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk-example/demo/button"
	"github.com/worldiety/wtk-example/demo/dialog"
	"github.com/worldiety/wtk-example/demo/index"
	"github.com/worldiety/wtk-example/demo/menu"
	"github.com/worldiety/wtk-example/demo/notfound"
	"github.com/worldiety/wtk-example/demo/textarea"
	"github.com/worldiety/wtk-example/demo/textfield"
	"github.com/worldiety/wtk-example/demo/typography"
	"github.com/worldiety/wtk/theme/material/icon"
	"strconv"
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

func (a *App) SetView(v View) {
	a.Application.SetView(NewFrame().SetView(v))
}

func (a *App) Start() {
	a.UnmatchedRoute(notfound.FromQuery)
	a.Route(index.Path, index.FromQuery)
	a.Route(button.Path, button.FromQuery)
	a.Route(typography.Path, typography.FromQuery)
	a.Route(textfield.Path, textfield.FromQuery)
	a.Route(textarea.Path, textarea.FromQuery)
	a.Route(dialog.Path, dialog.FromQuery)
	a.Route(menu.Path, menu.FromQuery)
	a.Application.Start()
}
