package view

import "github.com/golangee/forms-example/www/forms/dom"

// Renderable is a marker interface which identifies one of our three kinds of
// DOM manipulation primitives. This must be one of Component, Node or Modifier.
type Renderable interface {
	// nodeOrModifierOrComponent is our private marker contract.
	nodeOrModifierOrComponent()
}

// A Component interface is currently only implementable by embedding a View. You may ask "if there is just one
// implementation, why would you need an interface?". The answer is, "because you did not read carefully enough".
// There will be as many implementations, as you create, however, only a part for the Component contract can
// be introduced by embedding a View. An interface is required to rely on dynamic polymorphic method dispatching, which
// can only be achieved by using interfaces.
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
