package view

import "github.com/golangee/forms-example/www/dom"

// Renderable is a marker interface which identifies one of our three kinds of
// DOM manipulation primitives. This must be one of Component, Node or Modifier.
type Renderable interface {
	// nodeOrModifierOrComponent is our private marker contract.
	nodeOrModifierOrComponent()
}

// A Component interface is currently only implementable by embedding a View.
type Component interface {
	// Render returns a view root Node.
	Render() Node
	// Observe registers with the component which notifies for changes.
	Observe(f func()) Handle
	Renderable
}

// Render clears the body of the page and applies the given Renderable.
func RenderBody(c Renderable) {
	body := dom.GetWindow().Document().Body()
	body.Clear()
	WithElement(body, c).Element()
}
