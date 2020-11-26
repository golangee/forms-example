package base

import "github.com/golangee/forms-example/www/dom"

type RenderableView interface {
	Render() Renderable
	Observe(f func()) Handle
	RenderNode
}

type RenderNode interface {
	BuilderOrModifierOrRenderable()
}

type Renderable interface {
	CreateElement() dom.Element
	RenderNode
}

type BuilderFunc func() dom.Element

func (f BuilderFunc) CreateElement() dom.Element {
	return f()
}

func (f BuilderFunc) BuilderOrModifierOrRenderable() {
	panic("marker interface")
}

type Modifier interface {
	Modify(e dom.Element)
	RenderNode
}

type ModifierFunc func(e dom.Element)

func (f ModifierFunc) Modify(e dom.Element) {
	f(e)
}

func (f ModifierFunc) BuilderOrModifierOrRenderable() {
	panic("marker interface")
}
