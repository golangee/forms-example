package app

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	t01helloworld "github.com/golangee/forms-example/www/internal/01-tutorial-forms/01-chapter-essentials/01-section-setup/01-helloworld"
	t02hellohtml "github.com/golangee/forms-example/www/internal/01-tutorial-forms/01-chapter-essentials/02-section-hello/02-hellohtml"
	t03helloparam "github.com/golangee/forms-example/www/internal/01-tutorial-forms/01-chapter-essentials/02-section-hello/03-helloparam"
	t04helloproperty "github.com/golangee/forms-example/www/internal/01-tutorial-forms/01-chapter-essentials/03-section-component/04-helloproperty"
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

	a.router.AddRoute("/", a.apply(a.index)).
		AddRoute("/tutorials", a.apply(tutorialOverview)).
		AddRoute(t01helloworld.Path, a.apply(t01helloworld.FromQuery)).
		AddRoute(t02hellohtml.Path, a.apply(t02hellohtml.FromQuery)).
		AddRoute(t03helloparam.Path, a.apply(t03helloparam.FromQuery)).
		AddRoute(t04helloproperty.Path, a.apply(t04helloproperty.FromQuery))

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
