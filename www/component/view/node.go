package view

import "github.com/golangee/forms-example/www/dom"

// A Node gives access to a dom.Element. Usually the (render) Node allocates a new
// Element, which can be attached.
type Node interface {
	// Element returns the wrapped Element, usually a new instance is allocated.
	Element() dom.Element
	Renderable
}

// NodeFunc is a func type which allows usage as a Node. It is the only way to create such
// an implementation.
type NodeFunc func() dom.Element

// Element returns the wrapped Element, usually a new instance is allocated.
func (f NodeFunc) Element() dom.Element {
	return f()
}

// nodeOrModifierOrComponent is our private marker contract.
func (f NodeFunc) nodeOrModifierOrComponent() {
}
