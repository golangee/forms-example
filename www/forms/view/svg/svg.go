package svg

import (
	"github.com/golangee/forms-example/www/forms/dom"
	. "github.com/golangee/forms-example/www/forms/view"
)

func Svg(e ...Renderable) Node {
	return ElementNS("http://www.w3.org/2000/svg", "svg", e...)
}

func Fill(fill string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("fill", fill)
	})
}

func Stroke(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke", stroke)
	})
}

func StrokeLinecap(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-linecap", stroke)
	})
}

func StrokeLinejoin(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-linejoin", stroke)
	})
}

func StrokeWidth(stroke string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("stroke-width", stroke)
	})
}

func ViewBox(viewBox string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("viewBox", viewBox)
	})
}

func Path(e ...Renderable) Node {
	return ElementNS("http://www.w3.org/2000/svg", "path", e...)
}

func D(d string) Modifier {
	return ModifierFunc(func(e dom.Element) {
		e.SetAttribute("d", d)
	})
}
