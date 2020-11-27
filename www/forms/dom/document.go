package dom

import "syscall/js"

type Document struct {
	val js.Value
}

func newDocument(val js.Value) Document {
	return Document{val}
}

func (n Document) Body() Element {
	return newElement(n.val.Get("body"))
}

// In an HTML document, the document.createElement() method creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
func (n Document) createTextNode(name string) js.Value {
	return n.val.Call("createTextNode", name)
}

// In an HTML document, the document.createElement() method creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
func (n Document) CreateElement(name string) Element {
	v := n.val.Call("createElement", name)
	return newElement(v)
}

// In an HTML document, the document.createElement() method creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
func (n Document) CreateElementNS(ns string, name string) Element {
	v := n.val.Call("createElementNS", ns, name)
	return newElement(v)
}

func (n Document) DocumentElement() Element {
	body := n.val.Get("documentElement")
	return newElement(body)
}
