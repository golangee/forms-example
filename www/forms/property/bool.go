package property

import (
	"fmt"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

// Bool represents a typed and observable property.
type Bool struct {
	Property
}

func NewBool(value bool) *Bool {
	b := &Bool{}
	b.Set(value)
	return b
}

// Set updates the value and notifies each registered observer.
func (s *Bool) Set(v bool) {
	s.Property.Set(v)
}

// Get returns the current value.
func (s *Bool) Get() bool {
	if s.Property.Get() == nil {
		return false
	}

	return s.Property.Get().(bool)
}

// Bind reads the current value from dst into this value. However, every subsequent observed change of the
// property is written into dst.
func (s *Bool) Bind(dst *bool) Handle {

	h := s.Property.Observe(func(old, new interface{}) {
		*dst = new.(bool)
	})

	s.Set(*dst)
	return h
}

// Toggle inverts the current state.
func (s *Bool) Toggle() {
	newV := !s.Get()
	log.NewLogger().Print(ecs.Msg("new toggle value is " + fmt.Sprint(newV)))
	s.Set(newV)
}

// Observe registered a typed observer.
func (s *Bool) Observe(onDidSet func(old, new bool)) Handle {
	return s.Property.Observe(func(old, new interface{}) {
		if old == nil {
			old = false
		}

		if old != new {
			onDidSet(old.(bool), new.(bool))
		}
	})
}
