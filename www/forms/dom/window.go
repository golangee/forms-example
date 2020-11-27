package dom

import "syscall/js"

// A Releasable cleans up references and the resource must not be used afterwards anymore.
type Releasable interface {
	Release()
}

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

func (n Window) OnHashChange(f func()) Releasable {
	fun := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	n.val.Set("onhashchange", fun)

	return fun
}

func (n Window) HashChange(f func()) Releasable {
	fun := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	n.val.Set("hashchange", fun)

	return fun
}

func (n Window) OnPopState(f func()) Releasable {
	fun := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		return nil
	})
	n.val.Set("onpopstate", fun)

	return fun
}

func (n Window) Location() Location {
	return newLocation(n.val.Get("location"))
}

func (n Window) SetLocation(url string) Window {
	n.val.Set("location", url)
	return n
}
