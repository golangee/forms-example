package progress

import (
	"github.com/golangee/forms-example/www/forms/property"
	. "github.com/golangee/forms-example/www/forms/view"
	s "github.com/golangee/forms-example/www/forms/view/svg"
)

type Circle struct {
	View
	visible property.Bool
}

func NewInfiniteCircle() *Circle {
	c := &Circle{}
	c.visible.Attach(c)
	return c
}

func (c *Circle) VisibleProperty() *property.Bool {
	return &c.visible
}

func (c *Circle) Render() Node {
	return s.Svg(Class("text-primary stroke-current wtk-mdc-circular-progress"), s.ViewBox("25 25 50 50"),
		s.Circle(Class("wtk-mdc-circular-progress__path"), s.Cx("50"), s.Cy("50"), s.R("20"), s.Fill("none"), s.StrokeWidth("4"), s.StrokeMiterlimit("10")),
	)
}
