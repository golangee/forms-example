package layout

import . "github.com/golangee/forms-example/www/forms/view"

type VStack struct {
	View
	nodes []Renderable
}

func NewVStack(r ...Renderable) *VStack {
	return &VStack{
		nodes: r,
	}
}

func (c *VStack) Render() Node {
	return Div(Join(Class("grid grid-cols-1 justify-items-start"), c.nodes...)...)
}
