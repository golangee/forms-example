package base

import "github.com/golangee/forms-example/www/html"

// A Resource release method should be called for any resource, as soon as it is not required anymore to avoid
// memory leaks. Afterwards the Resource must not be used anymore.
// Even though we have a GC, we cannot rely on it, because the Resource may have registrations
// beyond our process, which requires holding global callback references, so that the outer system
// can call us. An example for this are go functions wrapped as callbacks in the wasm tier made available
// for the javascript DOM, like event handlers. Also cgo or rpc mechanism are possible.
type Resource interface {
	Release() // Release clean up references and the resource must not be used afterwards anymore.
}


// A View is a construction plan to describe how to build a view.
type View struct {
	stateful *Stateful
	Compose func()html.Renderable
}


func (v *View) getStateful() *Stateful {
	if v.stateful == nil {
		v.stateful = &Stateful{}
	}

	return v.stateful
}

func (v *View) Invalidate() {
	v.getStateful().Invalidate()
}

func (v *View) Observe(f func()) Handle {
	return v.getStateful().Observe(f)
}

/*
func (v *View) Render() html.Renderable {
	return html.RenderableFunc(
		func() dom.Element {
			var flatten []html.RenderableOrModifier
			for _, poser := range v.posers {
				flatten = append(flatten, poser()...)
			}

			r := flatten[0].(html.Renderable).Render()

			for i := 1; i < len(flatten); i++ {
				switch t := flatten[i].(type) {
				case html.Renderable:
					r.AppendElement( t.Render())
				case html.Modifier:
					t.Modify(r)
				default:
					panic(fmt.Sprint(t))
				}
			}
			return r
		},
	)
}*/

/*
func (v *View) Render() html.Renderable {
	log.NewLogger().Print(ecs.Msg("having " + strconv.Itoa(len(v.mod)) + "mods"))

	var firstCreator html.Renderable
	for _, modifier := range v.mod {
		if t, ok := modifier.(html.Renderable); ok {
			firstCreator = t
			break
		}
	}

	if firstCreator == nil {
		panic("compose must create a new element")
	}

	return html.RenderableFunc(func() dom.Element {
		e := firstCreator.Render()
		for _, modifier := range v.mod {
			switch t := modifier.(type) {
			case html.Modifier:
				t.Modify(e)
			}
		}

		return e
	})
}
*/
