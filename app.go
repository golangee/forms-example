package example

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
	"log"
	"strconv"
)

type App struct {
}

func NewApp() *App {
	return &App{}
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

func (a *App) Run() {
	log.Println("wasm done4")

	Root.RemoveAll()

	Run(Root, func() {
		myView := NewMyCustomComponent()
		//myView := section1.NewContentView()
		//myView := section2.NewContentView()
		//myView := section3.NewContentView()

		frame := NewFrame()
		frame.SetView(myView)
		Root.AddView(frame)
	})

	select {}
}
