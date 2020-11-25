package html

import "github.com/golangee/forms-example/www/dom"

type Composition []RenderableOrModifier

type RenderableOrModifier interface {
	renderableOrModifier()
}

type Renderable interface {
	Render() dom.Element
	RenderableOrModifier
}

type RenderableFunc func() dom.Element

func (f RenderableFunc) Render() dom.Element {
	return f()
}

func (f RenderableFunc) renderableOrModifier() {
	panic("marker interface")
}

type Modifier interface {
	Modify(e dom.Element)
	RenderableOrModifier
}

type ModifierFunc func(e dom.Element)

func (f ModifierFunc) Modify(e dom.Element) {
	f(e)
}

func (f ModifierFunc) renderableOrModifier() {
	panic("marker interface")
}
