package view

import (
	"fmt"
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/forms-example/www/forms/property"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"reflect"
	"strings"
)

func Element(name string, rm ...Renderable) Node {
	return WithElement(dom.GetWindow().Document().CreateElement(name), rm...)
}

func ElementNS(namespace, name string, rm ...Renderable) Node {
	return WithElement(dom.GetWindow().Document().CreateElementNS(namespace, name), rm...)
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
			case nil:
				// this makes optional sub-components easier
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

func Hr(e ...Renderable) Node {
	return Element("hr", e...)
}

func Button(e ...Renderable) Node {
	return Element("button", e...)
}

func Nav(e ...Renderable) Node {
	return Element("nav", e...)
}

// Yield is a convenience operator to apply or insert multiple renderables as one, especially useful if one
// ever needs to evaluate or append multiple renderables without a container.
func Yield(r ...Renderable) Renderable {
	return ModifierFunc(func(e dom.Element) {
		for _, renderable := range r {
			WithElement(e, renderable)
		}
	})
}

// Join is a convenience operator to merge 1+n renderables into a common slice.
func Join(r Renderable, other ...Renderable) []Renderable {
	tmp := make([]Renderable, 0, len(other)+1)
	tmp = append(tmp, r)
	for _, x := range other {
		tmp = append(tmp, x)
	}
	return tmp
}

// Class sets and replaces all existing class definitions.
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

// AddClass appends each given class.
func AddClass(classes ...string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if len(classes) == 0 {
			return
		}

		for _, class := range classes {
			if strings.ContainsRune(class, ' ') {
				// separate it
				for _, s := range strings.Split(class, " ") {
					e.AddClass(s)
				}
			} else {
				e.AddClass(class)
			}
		}
	})
}

// RemoveClass remove each given class.
func RemoveClass(classes ...string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if len(classes) == 0 {
			return
		}

		for _, class := range classes {
			if strings.ContainsRune(class, ' ') {
				// separate it
				for _, s := range strings.Split(class, " ") {
					e.RemoveClass(s)
				}
			} else {
				e.RemoveClass(class)
			}
		}
	})
}

func WithModifiers(m ...Modifier) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for _, modifier := range m {
			modifier.Modify(e)
		}
	})
}

// If applies the given positive and negative modifiers in-place, without causing
// an entire re-rendering, if the property changes. This improves performance
// a lot.
func If(p *property.Bool, pos, neg Modifier) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if p.Get() {
			if pos!=nil{
				pos.Modify(e)
			}
		} else {
			if neg!=nil{
				neg.Modify(e)
			}
		}

		h := p.Observe(func(old, new bool) {
			if new {
				pos.Modify(e)
			} else {
				neg.Modify(e)
			}
		})

		e.AddReleaseListener(h.Release)
	})
}

// Display sets the css style attribute "display".
func Display(display string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty("display", display)
	})
}

// Visibility sets the css style attribute "visibility".
// See also https://developer.mozilla.org/de/docs/Web/CSS/visibility
func Visibility(visibility string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty("visibility", visibility)
	})
}

// PointerEvents sets the CSS style attribute pointer-events,
// see also https://developer.mozilla.org/de/docs/Web/CSS/pointer-events.
func PointerEvents(visibility string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty("pointer-events", visibility)
	})
}

// BackgroundColor sets the CSS style attribute background-color,
// see also https://developer.mozilla.org/de/docs/Web/CSS/background-color.
func BackgroundColor(color string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty("background-color", color)
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

func TabIndex(t string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("tabindex", t)
	})
}

func I(mods ...Renderable) Node {
	return Element("i", mods...)
}

func A(mods ...Renderable) Node {
	return Element("a", mods...)
}

func Em(mods ...Renderable) Node {
	return Element("em", mods...)
}

func Alt(a string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("alt", a)
	})
}

func Title(a string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Set("title", a)
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

func Pre(mods ...Renderable) Node {
	return Element("pre", mods...)
}

func Code(mods ...Renderable) Node {
	return Element("code", mods...)
}

func Aside(mods ...Renderable) Node {
	return Element("aside", mods...)
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

func With(f func() Renderable) Modifier {
	return ModifierFunc(func(e dom.Element) {
		x := f()
		if x == nil {
			return
		}

		switch t := x.(type) {
		case Node:
			e.AppendElement(t.Element())
		case Modifier:
			t.Modify(e)
		default:
			panic(fmt.Sprint(e))
		}
	})
}

func InsideDom(f func(e dom.Element)) Modifier {
	return ModifierFunc(func(e dom.Element) {
		f(e)
	})
}

func ForEach(len int, f func(i int) Renderable) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for i := 0; i < len; i++ {
			x := f(i)
			WithElement(e, x).Element()
			/*switch t := x.(type) {
			case Node:
				e.AppendElement(t.Element())
			case Modifier:
				t.Modify(e)
			default:
				panic(fmt.Sprint(reflect.TypeOf(x), e))
			}*/
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

func AddKeyDownListener(f func(keyCode int)) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.AddKeyListener("keydown", f)
	})
}

func AddEventListenerOnce(eventType string, f func()) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.AddEventListener(eventType, true, f)
	})
}

func AriaLabel(label string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-label", label)
	})
}

func Role(role string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("role", role)
	})
}
