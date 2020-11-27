package dom

import "syscall/js"

type Global struct {
	val js.Value
}

func GetGlobal() Global {
	return Global{val: js.Global()}
}

func (n Global) Get(name string) Element {
	return newElement(n.val.Get(name))
}
