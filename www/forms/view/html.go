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
				for _, modifier := range t.getPostModifiers() {
					modifier.Modify(x)
				}
				var observer func()
				var xHandle Handle
				observer = func() {
					x.Release()
					xHandle.Release()
					newElem := t.Render().Element()
					for _, modifier := range t.getPostModifiers() {
						modifier.Modify(newElem)
					}
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

func IFrame(e ...Renderable) Node {
	return Element("iframe", e...)
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
			WithElement(e, renderable).Element()
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
					s = strings.TrimSpace(s)
					if s != "" {
						e.AddClass(s)
					}
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

func SetAttribute(attr, value string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute(attr, value)
	})
}

func RemoveAttribute(attr string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.RemoveAttribute(attr)
	})
}

func WithModifiers(m ...Modifier) Modifier {
	return ModifierFunc(func(e dom.Element) {
		for _, modifier := range m {
			modifier.Modify(e)
		}
	})
}

// If only evaluates the flag once and can not be changed afterwards. It is useful in non-components
// or if properties are not needed, because a full rendering will be done anyway. See also IfCond for
// a more efficient way of changing properties.
func If(flag bool, pos, neg Renderable) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if flag {
			if pos != nil {
				WithElement(e, pos).Element()
			}
		} else {
			if neg != nil {
				WithElement(e, neg).Element()
			}
		}
	})
}

// IfCond applies the given positive and negative modifiers in-place, without causing
// an entire re-rendering, if the property changes. This improves performance
// a lot. See also If.
func IfCond(p *property.Bool, pos, neg Modifier) Modifier {
	return ModifierFunc(func(e dom.Element) {
		if p.Get() {
			if pos != nil {
				pos.Modify(e)
			}
		} else {
			if neg != nil {
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

// Style sets a single CSS property.
func Style(property, value string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Style().SetProperty(property, value)
	})
}

func Text(t string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.AppendTextNode(t)
	})
}

func InnerHTML(t string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetInnerHTML(t)
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

//TODO useful?
func Focus() Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.Call("focus")
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

// WithComponent post-modifies the given Component for each future rendering.
func WithComponent(r Component, mods ...Modifier) Component {
	r.setPostModifiers(mods)
	return r
}

func ID(id string)Modifier{
	return ModifierFunc(func(e dom.Element) {
		e.SetID(id)
	})
}

// With post-modifies the given Renderable for each future rendering.
func With(r Renderable, mods ...Modifier) Renderable {
	switch t := r.(type) {
	case Component:
		return WithComponent(t, mods...)
	case Node:
		return NodeFunc(func() dom.Element {
			elem := t.Element()
			tmp := make([]Renderable, 0, len(mods))
			for _, mod := range mods {
				tmp = append(tmp, mod)
			}
			WithElement(elem, tmp...)
			return elem
		})
	case Modifier:
		return ModifierFunc(func(e dom.Element) {
			for _, mod := range mods {
				mod.Modify(e)
			}
		})
	default:
		panic(fmt.Sprintf(reflect.TypeOf(r).String()))
	}
}

/* TODO what is that use case?
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
}*/

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
		e.AddKeyListener("keyup", f)
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

func AriaOrientation(orientation string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-orientation", orientation)
	})
}

func AriaLabelledby(label string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("aria-labelledby", label)
	})
}

func Role(role string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("role", role)
	})
}
