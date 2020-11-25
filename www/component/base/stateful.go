package base

import "sync"

type Handle struct {
	parent *Stateful
	idx    int
}

func (h Handle) Release() {
	h.parent.lock.Lock()
	defer h.parent.lock.Unlock()

	h.parent.observers[h.idx] = nil
}

type Stateful struct {
	observers []func()   // we only append, invalid observers will be set to nil
	lock      sync.Mutex // we ever need this very short, no deadlocks possible
}

func (c *Stateful) Invalidate() {
	c.lock.Lock()
	length := len(c.observers)
	c.lock.Unlock()

	for i := 0; i < length; i++ {
		c.lock.Lock()
		observer := c.observers[i]
		c.lock.Unlock()

		if observer != nil {
			observer()
		}

	}
}

func (c *Stateful) Observe(f func()) Handle {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.observers = append(c.observers, f)
	return Handle{idx: len(c.observers)}
}
