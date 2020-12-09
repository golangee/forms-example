package modal

import (
	"github.com/golangee/forms-example/www/forms/dom"
	. "github.com/golangee/forms-example/www/forms/view"
	"strconv"
)

// An Overlay should be added to the body element and takes over the entire screen. It is fully transparent
// and allows to place multiple Renderable at arbitrary absolute positions. I know Popper
// (https://github.com/popperjs/popper-core#why-popper) but it is really hard to integrate into our lifecycle
// and I was not able to get it working in a reasonable time, so we drive our own implementation.
type Overlay struct {
	View
	items          []*overlayLayoutParams
	clickListeners []func()
}

func NewOverlay() *Overlay {
	c := &Overlay{}
	c.Observe(func() {
		c.alignOverlay()
	})
	return c
}

func (c *Overlay) AddClickListener(f func()) *Overlay {
	c.clickListeners = append(c.clickListeners, f)
	return c
}

func (c *Overlay) Put(id string, content Renderable) *Overlay {
	c.items = append(c.items, &overlayLayoutParams{
		domId:   id,
		content: content,
	})

	c.Invalidate()

	return c
}

func (c *Overlay) alignOverlay() {
	for _, item := range c.items {
		target := dom.GetDocument().GetElementById(item.domId)
		if !target.IsNull() && item.overlayElem != nil {
			overlayRect := item.overlayElem.GetBoundingClientRect()
			targetRect := target.GetBoundingClientRect()

			wndWidth := dom.GetWindow().InnerWidth()

			scrollbars := 4
			wantedLeft := targetRect.GetLeft()
			if overlayRect.GetWidth()+wantedLeft > wndWidth {
				wantedLeft = wndWidth - overlayRect.GetWidth() - scrollbars
			}

			if overlayRect.GetWidth() > wndWidth {
				overlayRect.SetWidth(wndWidth - scrollbars)
			}

			item.overlayElem.Style().SetProperty("left", strconv.Itoa(wantedLeft)+"px")
			item.overlayElem.Style().SetProperty("top", strconv.Itoa(targetRect.GetBottom())+"px")

		}
	}
}

func (c *Overlay) Render() Node {
	return Div(Class("absolute inset-0"),
		AddClickListener(func() {
			for _, listener := range c.clickListeners {
				listener()
			}
		}),
		ForEach(len(c.items), func(i int) Renderable {
			lp := c.items[i]
			return With(lp.content, Style("position", "absolute"), InsideDom(func(e dom.Element) {
				lp.overlayElem = &e
				c.alignOverlay()
			}))
		}),
	)
}

type overlayLayoutParams struct {
	domId       string
	content     Renderable
	overlayElem *dom.Element
}

// ShowOverlay displays the overlay in the body element and removes it when clicked.
func ShowOverlay(overlay *Overlay) {
	body := dom.GetWindow().Document().Body()
	resizeListener := dom.GetWindow().AddEventListener("resize", overlay.alignOverlay)
	WithElement(body, overlay).Element()
	overlayElem := body.LastChild()
	overlay.AddClickListener(func() {
		overlayElem.Remove()
		resizeListener.Release()
	})
	dom.Post(overlay.alignOverlay)
}
