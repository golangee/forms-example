package view

import "github.com/golangee/forms-example/www/dom"

// Modifier is Renderable which changes attributes or contents of the given dom.Element.
type Modifier interface {
	// Modify applies its changes to the given element.
	Modify(e dom.Element)
	Renderable
}

// ModifierFunc is a func type which allows usage as a Modifier. It is the only way to create such
// an implementation.
type ModifierFunc func(e dom.Element)

// Modify applies its changes to the given element.
func (f ModifierFunc) Modify(e dom.Element) {
	f(e)
}

// nodeOrModifierOrComponent is our private marker contract.
func (f ModifierFunc) nodeOrModifierOrComponent() {
}
