package highlightjs

import (
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/forms-example/www/forms/property"
	. "github.com/golangee/forms-example/www/forms/view"
)

type CodeView struct {
	code property.String
	lang property.String
	View
}

func NewCode() *CodeView {
	c := &CodeView{}
	c.code.Attach(c)
	return c
}

func (c *CodeView) CodeProperty() *property.String {
	return &c.code
}

func (c *CodeView) LangProperty() *property.String {
	return &c.lang
}

func (c *CodeView) Render() Node {
	return Pre(Class(c.lang.Get()),
		Code(
			Text(c.code.Get()),
		),
		// order is important here, because otherwise round the Code block has not been created yet
		InsideDom(func(e dom.Element) {

			dom.GetGlobal().Get("hljs").Call("highlightBlock", e)
			dom.GetGlobal().Get("hljs").Call("lineNumbersBlock", e)
		}),
	)
}
