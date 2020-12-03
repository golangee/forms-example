package app

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/internal/index"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

type Application struct {
	router *router.Router
	log    log.Logger
}

func NewApplication() *Application {
	a := &Application{
		router: router.NewRouter(),
		log:    log.NewLogger(ecs.Log("application")),
	}

	a.router.AddRoute("/", a.apply(tutorialOverview))

	for _, chapter := range index.Tutorial.Fragments {
		for _, section := range chapter.Fragments {
			path := "/" + index.Tutorial.ID() + "/" + chapter.ID() + "/" + section.ID()
			a.router.AddRoute(path, a.apply(tutorialStepview))
		}
	}

	a.router.
		SetUnhandledRouteAction(a.apply(func(query router.Query) Renderable {
			return Span(Text("unmatched route to " + query.Path()))
		}))

	return a
}

func (a *Application) apply(f func(query router.Query) Renderable) func(query router.Query) {
	return func(query router.Query) {
		RenderBody(a.page(query, f(query)))
	}
}

func (a *Application) index(q router.Query) Renderable {
	return P(Text("Welcome to the forms tutorial"))
}

func (a *Application) Run() {
	a.router.Start()
	select {}
}
