// Package popper wraps a small javascript library which solves a bunch of pop-over problems. See
// https://github.com/popperjs/popper-core#why-popper for details. You need to include
// the javascript popper library in your root index.gohtml template.
package popper

import (
	"fmt"
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/forms-example/www/forms/property"
	"github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

// Popper takes ownership over an anchor and a popover element.
type Popper struct {
	view.View
	anchor, popover view.Component
	isShown         property.Bool
}

// NewPopper creates a Popper instance and takes ownership over an anchor and a popover element.
func NewPopper(anchor, popover view.Component) *Popper {
	c := &Popper{
		anchor:  anchor,
		popover: popover,
	}

	return c
}

func (c *Popper) ShowProperty() *property.Bool {
	return &c.isShown
}

func (c *Popper) Show() *Popper {
	c.isShown.Set(true)
	return c
}

func (c *Popper) Hide() *Popper {
	c.isShown.Set(false)
	return c
}

// Self assigns the receiver to the given pointer to reference.
func (c *Popper) Self(ref **Popper) *Popper {
	*ref = c
	return c
}

func (c *Popper) Render() view.Node {
	// This implementation is nothing I'm really happy about. It circumvents the entire
	// lifecycle system, because we need to attach a foreign JS lifecycle into our
	// system. By the way, it is a real certificate of poverty for the entire html5 world,
	// that one needs an 18KiB Javascript file to get working popovers. And yes, they are
	// right by saying that the usual CSS and Javascript example do not work properly.

	var anchorElem, popElem *dom.Element
	var popperInstance *dom.Element

	
	return view.Div(
		view.With(c.popover,
			view.IfCond(&c.isShown, view.Style("display", "block"), view.Style("display", "none")),
			view.InsideDom(func(e dom.Element) {
				popElem = &e
				log.NewLogger().Print(ecs.Msg("da popover: " + fmt.Sprint(popElem)))
			})),

		view.With(c.anchor,
			view.InsideDom(func(e dom.Element) {
				anchorElem = &e
				log.NewLogger().Print(ecs.Msg("createPopper" + fmt.Sprint(anchorElem) + "=>" + fmt.Sprint(popElem)))
				inst := dom.GetGlobal().Get("Popper").Call("createPopper", *anchorElem, *popElem)
				popperInstance = &inst
				anchorElem.AddReleaseListener(func() {
					log.NewLogger().Print(ecs.Msg("popper released"))
					if popperInstance != nil {
						popperInstance.Call("destroy")
						popperInstance = nil
						anchorElem = nil
						popElem = nil
					}
				})
			})),
	)
}
