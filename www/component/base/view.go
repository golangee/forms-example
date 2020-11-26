package base

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
	stateful   *Stateful
	Render    func() Renderable
	//activeElem *dom.Element
	tag        string
}

func (v *View)	BuilderOrModifierOrRenderable(){

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
func (v *View) Attach(r html.Renderable) html.Renderable {
	return html.BuilderFunc(func() dom.Element {
		elem := r.CreateElement()
		v.activeElem = &elem
		_= v.Observe(func() {
			newElem := r.CreateElement()

			log.NewLogger().Print(ecs.Msg(newElem.String()))
			log.NewLogger().Print(ecs.Msg("invalidate: "+v.String()))

			if v.activeElem != nil {
				v.activeElem.ReplaceWith(newElem)
				log.NewLogger().Print(ecs.Msg("replaced node: "+v.String()))
			}

			v.activeElem = &newElem

		})


		newElem.AddReleaseListener(func() {
			handle.Release()
		})

		return elem
	})
}*/

func (v *View) SetTag(tag string) {
	v.tag = tag
}

func (v *View) String() string {
	return "View-" + v.tag
}

/*
func (v *View) CreateElement() html.Renderable {
	return html.BuilderFunc(
		func() dom.Element {
			var flatten []html.RenderNode
			for _, poser := range v.posers {
				flatten = append(flatten, poser()...)
			}

			r := flatten[0].(html.Renderable).CreateElement()

			for i := 1; i < len(flatten); i++ {
				switch t := flatten[i].(type) {
				case html.Renderable:
					r.AppendElement( t.CreateElement())
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
func (v *View) CreateElement() html.Renderable {
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

	return html.BuilderFunc(func() dom.Element {
		e := firstCreator.CreateElement()
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
