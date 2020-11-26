package text

import (
	h "github.com/golangee/forms-example/www/component/base"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"strconv"
	"time"
)

type Text struct {
	text string
	h.View
}

func NewText(text string) *Text {
	c := &Text{}
	c.text = text
	c.SetTag("Text@"+strconv.Itoa(time.Now().Second()))

	var bla h.RenderableView
	bla = c
	_ = bla
	return c
}

func (c *Text) Render() h.Renderable {
	return h.Span(h.Text(c.text), h.AddEventListener("click", func() {
		c.SetText("hey " + time.Now().String())
		log.NewLogger().Print(ecs.Msg(c.text))
	}))
}

func (c *Text) SetText(text string) {
	c.text = text
	c.Invalidate()
}
