package dom

import (
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"syscall/js"
)

const EventRelease = "forms-release"

// Element is always used in a value context, to avoid additional GC pressure.
// Therefore our state-handling relies on Javascript messaging (especially the
// release cycles).
type Element struct {
	val js.Value
}

func newElement(val js.Value) Element {
	return Element{val}
}

func (n Element) SetTextContent(v string) Element {
	n.val.Set("textContent", v)
	return n
}

func (n Element) AppendTextNode(t string) Element {
	tn := GetWindow().Document().createTextNode(t)
	n.val.Call("appendChild", tn)
	return n
}

func (n Element) GetTagName() string {
	return n.val.Get("tagName").String()
}

func (n Element) SetClassName(str string) Element {
	n.val.Set("className", str)
	return n
}

func (n Element) AppendElement(aChild Element) Element {
	n.val.Call("appendChild", aChild.val)
	return n
}

func (n Element) Clear() Element {
	for _, element := range n.Children() {
		element.Release()
	}
	return n.SetTextContent("")
}

func (n Element) Children() []Element {
	var res []Element
	arr := n.val.Get("children")
	for i := 0; i < arr.Length(); i++ {
		res = append(res, newElement(arr.Index(i)))
	}
	return res
}

func (n Element) ReplaceWith(o Element) Element {
	n.val.Call("replaceWith", o.val)
	//	n.AppendElement(o)
	return o
}

func (n Element) String() string {
	return n.val.Get("outerHTML").String()
}

// Set sets the javascript property
func (n Element) Set(p string, x interface{}) Element {
	n.val.Set(p, x)
	return n
}

// AddEventListener is internally very complex, because it keeps a global callback
// reference to connect the wasm and the javascript context. The wasm side must keep
// a global un-collectable function and the javascript side does the same. This makes
// event handling currently very expensive. Always ensure that
// you call Release on this Element to free all resources.
func (n Element) AddEventListener(typ string, once bool, listener func()) Element {
	log.NewLogger().Print(ecs.Msg("addEventListener " + typ + " " + n.GetTagName()))

	defer GlobalPanicHandler()

	alreadyReleased := false
	var actualFunc, releaseFunc js.Func

	unregisterJS := func() {
		defer GlobalPanicHandler()

		if !alreadyReleased {
			alreadyReleased = true
			n.val.Call("removeEventListener", typ, actualFunc, once)
			n.val.Call("removeEventListener", EventRelease, releaseFunc, true)
			actualFunc.Release()
			releaseFunc.Release()
		}
	}

	actualFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer GlobalPanicHandler()

		listener()
		if once && !alreadyReleased {
			unregisterJS()
		}
		return nil
	})

	releaseFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer GlobalPanicHandler()

		log.NewLogger().Print(ecs.Msg("func: received release event"))
		if !alreadyReleased {
			unregisterJS()
		}
		return nil
	})

	n.val.Call("addEventListener", EventRelease, releaseFunc, true)
	n.val.Call("addEventListener", typ, actualFunc, once)

	return n
}

// Release is part of our custom lifecycle. We need a manual destructor, because we
// have currently two running contexts: our wasm program and the browsers javascript interpreter.
// This is important to remove callbacks and other attached resources, which would otherwise
// never be freed, due to global or cyclic references. All contained children element, will
// also receive a Release call.
func (n Element) Release() {
	for _, element := range n.Children() {
		element.Release()
	}

	event := js.Global().Get("Event").New(EventRelease)
	n.val.Call("dispatchEvent", event)
}

func (n Element) AddReleaseListener(f func()) Element {
	var fun js.Func
	fun = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer GlobalPanicHandler()

		f()

		// TODO not sure if 'once' is used correctly, but without removing, we get weired warnings
		n.val.Call("removeEventListener", EventRelease, fun, true)
		fun.Release()
		return nil
	})

	n.val.Call("addEventListener", EventRelease, fun, true)
	return n
}

func (n Element) AddClass(v string) Element {
	n.val.Get("classList").Call("add", v)
	return n
}

func (n Element) HasClass(v string) bool {
	return n.val.Get("classList").Call("contains", v).Bool()
}

func (n Element) RemoveClass(v string) Element {
	n.val.Get("classList").Call("remove", v)
	return n
}

// Equal has the Javascript == semantic on the Element (equal reference)
func (n Element) Equal(o Element) bool {
	return n.val.Equal(o.val)
}
