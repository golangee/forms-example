package example

import (
	"github.com/worldiety/wtk"
	"log"
	"strconv"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

type MyCustomComponent struct {
	text1   *wtk.Text
	text2   *wtk.Text
	counter int
}

func NewMyCustomComponent() *MyCustomComponent {
	c := &MyCustomComponent{}

	vstack := &wtk.VStack{}

	c.text1 = &wtk.Text{}
	c.text1.Value.Set("hello world")
	vstack.AddView(c.text1)

	c.text2 = &wtk.Text{}
	c.text2.Value.Set("a second text line")
	vstack.AddView(c.text2)

	btn := &wtk.Button{Text: "press me"}
	btn.AddOnClickListener(func() {
		c.counter++
		c.text2.Value.Set("pressed " + strconv.Itoa(c.counter))
	})
	vstack.AddView(btn)
	return c
}

func (a *App) Run() {
	log.Println("wasm done4")

	wtk.Root.RemoveAll()

	myView := NewMyCustomComponent()
	wtk.Root.AddView(myView)

	select {}
}
