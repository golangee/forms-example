package text

import (
	"github.com/golangee/forms-example/www/component/base"
	h "github.com/golangee/forms-example/www/html"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"time"
)

type Text struct {
	text string
	base.View
}

func NewText(text string) *Text {
	c := &Text{}
	c.text = text

	c.Describe(func() h.Composition {
		return h.Composition{
			h.Span(h.Text(c.text), h.DebugLog("compose: NewText"), h.AddEventListener("click", func() {
				c.SetText("hey " + time.Now().String() + " ->" + c.text)
				log.NewLogger().Print(ecs.Msg("hey"))
			})),
		}
	})

	return c
}

func (c *Text) SetText(text string) {
	c.text = text
	c.Invalidate()
}