package html

import "github.com/golangee/forms-example/www/dom"

type baseElement struct {
	dom.Element
}

func createBaseElement(name string) baseElement {
	return baseElement{dom.GetWindow().Document().CreateElement(name)}
}

type Modifier func(e dom.Element)


