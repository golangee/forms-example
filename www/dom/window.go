package dom

import "syscall/js"

type Window struct {
	val js.Value
}

func newWindow() Window {
	n := js.Global().Get("window")
	return Window{n}
}

func GetWindow() Window {
	return newWindow()
}

// Document returns a reference to the document contained in the window.
func (n Window) Document() Document {
	return newDocument(n.val.Get("document"))
}
