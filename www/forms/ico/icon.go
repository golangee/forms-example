package ico

import (
	. "github.com/golangee/forms-example/www/forms/view"
)

type Icon struct {
	name string
	View
}

func NewIcon(name string) *Icon {
	return &Icon{name: name}
}

func (c *Icon) Render() Node {
	return I(Class("material-icons align-sub"), Text(c.name))
}
