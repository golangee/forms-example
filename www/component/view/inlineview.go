package view

// InlineView is a helper to define views in an abstract way, without the
// need to declare a (private) struct with the according method.
type InlineView struct {
	View
	f func() Node
}

// NewInlineView creates an InlineView which invokes the given closure on Render.
func NewInlineView(render func() Node) *InlineView {
	return &InlineView{f: render}
}

// Render returns a view root Node.
func (v InlineView) Render() Node {
	return v.f()
}
