package modal

import (
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/forms-example/www/forms/property"
	. "github.com/golangee/forms-example/www/forms/view"
	"time"
)

type Dialog struct {
	isOpen  property.Bool
	content Renderable
	View
}

func NewDialog() *Dialog {
	c := &Dialog{}
	// we do not attach the property, because we don't want a full re-rendering
	return c
}

func (c *Dialog) Show() *Dialog {
	c.isOpen.Set(true)
	return c
}

func (c *Dialog) Hide() *Dialog {
	c.isOpen.Set(false)
	return c
}

func (c *Dialog) SetContent(r Renderable) *Dialog {
	c.content = r
	return c
}

func (c *Dialog) Render() Node {
	// semi-transparent content blocking layer
	return Div(
		Class(" z-10 fixed ease-in-out inset-0 bg-black opacity-0 transition-all duration-500"),

		If(&c.isOpen,
			WithModifiers(
				Style("visibility", "visible"),
				AddClass("opacity-50"),
			),
			WithModifiers(
				Style("visibility", "hidden"),
				RemoveClass("opacity-50"),
			),
		),

		Div(
			Class("absolute inset-0"),
			AddClickListener(c.isOpen.Toggle),

			c.content,
		),
	)
}

func (c *Dialog) AddCloseListener(f func()) *Dialog {
	c.isOpen.Observe(func(old, new bool) {
		if new == false {
			f()
		}
	})

	return c
}

// ShowDialog creates a new full screen plane and appends it to the body element. TODO onclose: do we remove it from body?
func ShowDialog(content Renderable) *Dialog {
	dlg := NewDialog()
	dlg.SetContent(content)
	body := dom.GetWindow().Document().Body()
	WithElement(body, dlg).Element()
	dlgElem := body.LastChild()
	dom.Post(func() {
		dlg.Show()
	})

	dlg.AddCloseListener(func() {
		dom.SetTimeout(500*time.Millisecond, func() {
			dlgElem.Remove()
		})
	})

	return dlg
}
