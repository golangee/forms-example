package view

import (
	"fmt"
	"github.com/golangee/forms-example/www/dom"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"strings"
)

func Element(name string, rm ...Renderable) Node {
	return WithElement(dom.GetWindow().Document().CreateElement(name), rm...)
}

func WithElement(elem dom.Element, rm ...Renderable) Node {
	return NodeFunc(func() dom.Element {

		for _, e := range rm {
			switch t := e.(type) {
			case Node:
				elem.AppendElement(t.Element())
			case Modifier:
				t.Modify(elem)
			case Component:
				x := t.Render().Element()
				var observer func()
				var xHandle Handle
				observer = func() {
					x.Release()
					xHandle.Release()
					newElem := t.Render().Element()
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

func Div(e ...Renderable) Node {
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
		//e.SetTextContent(t)
		e.AppendTextNode(t)
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

func Figure(mods ...Renderable) Node {
	return Element("figure", mods...)
}

func Ul(mods ...Renderable) Node {
	return Element("ul", mods...)
}

func Li(mods ...Renderable) Node {
	return Element("li", mods...)
}

func Ol(mods ...Renderable) Node {
	return Element("ol", mods...)
}

func Img(mods ...Renderable) Node {
	return Element("img", mods...)
}

func P(mods ...Renderable) Node {
	return Element("p", mods...)
}

func Blockquote(mods ...Renderable) Node {
	return Element("blockquote", mods...)
}

func Figcaption(mods ...Renderable) Node {
	return Element("figcaption", mods...)
}

func Span(mods ...Renderable) Node {
	return Element("span", mods...)
}

func ForEach(len int, f func(i int) Renderable) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for i := 0; i < len; i++ {
			x := f(i)
			switch t := x.(type) {
			case Node:
				e.AppendElement(t.Element())
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

func AddClickListener(f func()) Modifier {
	return AddEventListener("click", f)
}

func AddEventListenerOnce(eventType string, f func()) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, true, f)
	})
}
