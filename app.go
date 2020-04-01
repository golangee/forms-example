package example

import (
	. "github.com/worldiety/wtk"
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
		NewText("hello world 2").Self(&c.text1),
		NewText("a second text line").Self(&c.text2),
		NewButton("press me").AddClickListener(func(v View) {
			c.counter++
			c.text1.SetValue("text 1: clicked " + strconv.Itoa(c.counter))
			c.text2.SetValue("text 2: clicked " + strconv.Itoa(c.counter))

			text := NewText("your click no " + strconv.Itoa(c.counter))
			c.VStack.AddViews(text)
		}),
	).Self(&c.VStack)

	return c
}

func (a *App) Run() {
	log.Println("wasm done4")

	Root.RemoveAll()

	myView := NewMyCustomComponent()
	Root.AddView(myView)

	select {}
}
