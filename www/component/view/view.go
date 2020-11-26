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
type View struct {
	stateful *Stateful
	//activeElem *dom.Element
	tag string
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

//deprecated: TODO only a debug feature?
func (v *View) SetTag(tag string) {
	v.tag = tag
}

// String returns a human debuggable name.
func (v *View) String() string {
	return "View-" + v.tag
}
