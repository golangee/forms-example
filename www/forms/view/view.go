package view

// A Resource release method should be called for any resource, as soon as it is not required anymore to avoid
// memory leaks. Afterwards the Resource must not be used anymore.
// Even though we have a GC, we cannot rely on it, because the Resource may have registrations
// beyond our process, which requires holding global callback references, so that the outer system
// can call us. An example for this are go functions wrapped as callbacks in the wasm tier made available
// for the javascript DOM, like event handlers. Also cgo or rpc mechanism are possible.
type Resource interface {
	Release() // Release clean up references and the resource must not be used afterwards anymore.
}

// A View is a construction plan to describe how to build a view.
// Intentionally, this View does not provide a Render()Node method for two important reasons:
//  1. a View for it alone would be a Component, however without any useable benefit
//  2. because there is no overloading, we cannot give a hint, to implement it correctly
type View struct {
	stateful      *Stateful
	postModifiers []Modifier
}

// nodeOrModifierOrComponent is our private marker contract.
func (v *View) nodeOrModifierOrComponent() {
}

// getStateful ensures a stateful instance
func (v *View) getStateful() *Stateful {
	if v.stateful == nil {
		v.stateful = &Stateful{}
	}

	return v.stateful
}

// Invalidate notifies all registered observers. Call it, to trigger a new
// render cycle for your Component.
func (v *View) Invalidate() {
	v.getStateful().Invalidate()
}

// Observe registers a callback which is invoked, when Invalidate has been called.
func (v *View) Observe(f func()) Handle {
	return v.getStateful().Observe(f)
}

func (v *View) getPostModifiers() []Modifier {
	return v.postModifiers
}

func (v *View) setPostModifiers(mods []Modifier) {
	v.postModifiers = mods
}
