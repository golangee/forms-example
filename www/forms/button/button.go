package button

import (
	"github.com/golangee/forms-example/www/forms/ico"
	v "github.com/golangee/forms-example/www/forms/view"
)

// A Button in the meaning of a material design text button.
// See also https://material.io/components/buttons#anatomy.
type Button struct {
	v.View
	content v.Renderable
	action  func()
}

func NewButton(action func()) *Button {
	c := &Button{
		action: action,
	}
	return c
}

func (c *Button) SetContent(content v.Renderable) *Button {
	c.content = content
	c.Invalidate()
	return c
}

func (c *Button) Render() v.Node {
	return v.Button(
		v.Style("min-width", "64px"),
		v.Class("text-left hover:bg-primary bg-opacity-10 bg-transparent text-primary focus:outline-none p-2 rounded text-center"), // not w-full
		v.AddClickListener(c.action),
		c.content,
	)
}

func NewIconTextButton(icon string, text string, action func()) *Button {
	return NewButton(action).SetContent(v.Span(ico.NewIcon(icon), v.Span(v.Class("pr-2")), v.Text(text)))
}

func NewTextButton(text string, action func()) *Button {
	return NewButton(action).SetContent(v.Span(v.Text(text)))
}
