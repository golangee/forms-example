package button

import (
	"github.com/golangee/forms-example/www/forms/ico"
	. "github.com/golangee/forms-example/www/forms/view"
)

// A TextButton in the meaning of a material design text button.
// See also https://material.io/components/buttons#anatomy.
type TextButton struct {
	View
	caption Renderable
	action  func()
}

func NewTextButton(action func()) *TextButton {
	c := &TextButton{
		action: action,
	}
	return c
}

func (c *TextButton) SetCaption(caption Renderable) *TextButton {
	c.caption = caption
	c.Invalidate()
	return c
}

func (c *TextButton) Render() Node {
	return Button(
		//	Class(style.HoverBgPrimary, style.HoverTextOnPrimary, style.BgTransparent, style.TextPrimary, "focus:outline-none"),
		Class("w-full text-left hover:bg-primary bg-opacity-10 bg-transparent text-primary focus:outline-none pt-2 pb-2 pl-3 pr-3 rounded"),
		AddClickListener(c.action),
		c.caption,
	)
}

func NewIconTextButton(icon string, text string, action func()) Renderable {
	return NewTextButton(action).SetCaption(Span(ico.NewIcon(icon), Span(Class("pr-2")), Text(text)))
}
