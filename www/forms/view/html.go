package view

import (
	"fmt"
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"reflect"
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
				panic(fmt.Sprintf("the type '%s' must be either a Node, a Modifier or a Component. Did you forget to add the Render method?", reflect.TypeOf(t).String()))
			}

		}

		return elem
	})
}

func Div(e ...Renderable) Node {
	return Element("div", e...)
}

func Nav(e ...Renderable) Node {
	return Element("nav", e...)
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

func I(mods ...Renderable) Node {
	return Element("i", mods...)
}

func A(mods ...Renderable) Node {
	return Element("a", mods...)
}

func Alt(a string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("alt", a)
	})
}

func Href(href string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("href", href)
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
