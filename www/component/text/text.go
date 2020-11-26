package text

import (
	v "github.com/golangee/forms-example/www/component/view"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"strconv"
	"time"
)

type Text struct {
	text string
	v.View
}

func NewText(text string) *Text {
	c := &Text{}
	c.text = text
	c.SetTag("Text@"+strconv.Itoa(time.Now().Second()))

	var bla v.Component
	bla = c
	_ = bla
	return c
}

func (c *Text) Render() v.Node {
	return v.Span(v.Text(c.text), v.AddEventListener("click", func() {
		c.SetText("hey " + time.Now().String())
		log.NewLogger().Print(ecs.Msg(c.text))
	}))
}

func (c *Text) SetText(text string) {
	c.text = text
	c.Invalidate()
}
