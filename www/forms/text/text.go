package text

import (
	v "github.com/golangee/forms-example/www/forms/view"
)

type Text struct {
	text string
	rs   []v.Renderable
	v.View
}

func NewText(text string, rs ...v.Renderable) *Text {
	return &Text{
		text: text,
		rs:   rs,
	}
}

func (c *Text) Render() v.Node {
	return v.Span(append([]v.Renderable{v.Text(c.text)}, c.rs...)...)
}

func (c *Text) SetText(text string) {
	c.text = text
	c.Invalidate()
}

func (c *Text) Text() string {
	return c.text
}
