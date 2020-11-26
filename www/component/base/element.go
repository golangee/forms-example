package base

import (
	"fmt"
	"github.com/golangee/forms-example/www/dom"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"strings"
)

func Element(name string, rm ...RenderNode) Renderable {
	return BuilderFunc(func() dom.Element {
		elem := dom.GetWindow().Document().CreateElement(name)

		for _, e := range rm {
			switch t := e.(type) {
			case Renderable:
				elem.AppendElement(t.CreateElement())
			case Modifier:
				t.Modify(elem)
			case RenderableView:
				x := t.Render().CreateElement()
				var observer func()
				var xHandle Handle
				observer = func() {
					x.Release()
					xHandle.Release()
					newElem := t.Render().CreateElement()
					x = x.ReplaceWith(newElem)
					xHandle = t.Observe(observer)
					x.AddReleaseListener(func() {
						xHandle.Release()
					})
				}
				xHandle = t.Observe(observer)

				x.AddReleaseListener(func() {
					xHandle.Release()
				})
				elem.AppendElement(x)
			default:
				panic(fmt.Sprint(e))
			}

		}

		return elem
	})
}

func Div(e ...RenderNode) Renderable {
	return Element("div", e...)
}

func Class(classes ...string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if len(classes) == 0 {
			return
		}

		if len(classes) == 1 {
			e.SetClassName(classes[0])
			return
		}

		e.SetClassName(strings.Join(classes, " "))
	})
}

func Text(t string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetTextContent(t)
	})
}

func Src(src string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("src", src)
	})

}

func Alt(a string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("alt", a)
	})
}

func Width(w string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("width", w)
	})
}

func Height(h string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("height", h)
	})
}

func Figure(mods ...RenderNode) Renderable {
	return Element("figure", mods...)
}

func Ul(mods ...RenderNode) Renderable {
	return Element("ul", mods...)
}

func Li(mods ...RenderNode) Renderable {
	return Element("li", mods...)
}

func Ol(mods ...RenderNode) Renderable {
	return Element("ol", mods...)
}

func Img(mods ...RenderNode) Renderable {
	return Element("img", mods...)
}

func P(mods ...RenderNode) Renderable {
	return Element("p", mods...)
}

func Blockquote(mods ...RenderNode) Renderable {
	return Element("blockquote", mods...)
}

func Figcaption(mods ...RenderNode) Renderable {
	return Element("figcaption", mods...)
}

func Span(mods ...RenderNode) Renderable {
	return Element("span", mods...)
}

func ForEach(len int, f func(i int) RenderNode) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for i := 0; i < len; i++ {
			x := f(i)
			switch t := x.(type) {
			case Renderable:
				e.AppendElement(t.CreateElement())
			case Modifier:
				t.Modify(e)
			default:
				panic(fmt.Sprint(e))
			}
		}
	})
}

func DebugLog(msg string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		log.NewLogger().Print(ecs.Msg(msg), ecs.ErrStack())

	})
}

func AddEventListener(eventType string, f func()) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, false, f)
	})
}

func AddEventListenerOnce(eventType string, f func()) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, true, f)
	})
}
