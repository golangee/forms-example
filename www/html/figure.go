package html

import "github.com/golangee/forms-example/www/dom"

func Class(classes string) Modifier {
	return func(e dom.Element) {
		e.SetClassName(classes)
	}
}

func Src(src string) Modifier {
	return func(e dom.Element) {
		e.Set("src", src)
	}
}

func Alt(a string) Modifier {
	return func(e dom.Element) {
		e.Set("alt", a)
	}
}

func Width(w string) Modifier {
	return func(e dom.Element) {
		e.Set("width", w)
	}
}

func Height(h string) Modifier {
	return func(e dom.Element) {
		e.Set("height", h)
	}
}

func Text(t string) Modifier {
	return func(e dom.Element) {
		e.Set("textContent", t)
	}
}

func Append(elems ...dom.Element) Modifier {
	return func(e dom.Element) {
		for _, elem := range elems {
			e.AppendElement(elem)
		}
	}
}

func createAndAppendElem(name string, mods ...Modifier) Modifier {
	e := dom.GetWindow().Document().CreateElement(name)

	for _, mod := range mods {
		mod(e)
	}

	return func(o dom.Element) {
		o.AppendElement(e)
	}
}

func Figure(mods ...Modifier) Modifier {
	return createAndAppendElem("figure", mods...)
}

func Div(mods ...Modifier) Modifier {
	return createAndAppendElem("div", mods...)
}

func Ul(mods ...Modifier) Modifier {
	return createAndAppendElem("ul", mods...)
}

func Li(mods ...Modifier) Modifier {
	return createAndAppendElem("li", mods...)
}

func Ol(mods ...Modifier) Modifier {
	return createAndAppendElem("ol", mods...)
}

func Img(mods ...Modifier) Modifier {
	return createAndAppendElem("img", mods...)
}

func P(mods ...Modifier) Modifier {
	return createAndAppendElem("p", mods...)
}

func Blockquote(mods ...Modifier) Modifier {
	return createAndAppendElem("blockquote", mods...)
}

func Figcaption(mods ...Modifier) Modifier {
	return createAndAppendElem("figcaption", mods...)
}

func Self(e *dom.Element) Modifier {
	return func(o dom.Element) {
		*e = o
	}
}

func ForEach(len int, f func(i int, e dom.Element)) Modifier {
	return func(e dom.Element) {
		for i := 0; i < len; i++ {
			f(i, e)
		}
	}
}

func AddEventListener(eventType string, once bool, f func(e dom.Element)) Modifier {
	return func(e dom.Element) {
		e.AddEventListener(eventType, once, f)
	}
}
