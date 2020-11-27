package app

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/internal/tutorial/01-helloworld"
	"github.com/golangee/forms-example/www/internal/tutorial/02-hellohtml"
	"github.com/golangee/forms-example/www/internal/tutorial/03-helloparam"
	 "github.com/golangee/forms-example/www/internal/tutorial/04-helloproperty"
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

	a.router.AddRoute("/", apply(a.index)).
		AddRoute(t01helloworld.Path, apply(t01helloworld.FromQuery)).
		AddRoute(t02hellohtml.Path, apply(t02hellohtml.FromQuery)).
		AddRoute(t03helloparam.Path, apply(t03helloparam.FromQuery)).
		AddRoute(t04helloproperty.Path, apply(t04helloproperty.FromQuery))

	a.router.
		SetUnhandledRouteAction(apply(func(query router.Query) Renderable {
			return Div(
				Text("unmatched route: "+query.Path()),
				a.index(query),
			)
		}))

	return a
}

func apply(f func(query router.Query) Renderable) func(query router.Query) {
	return func(query router.Query) {
		RenderBody(f(query))
	}
}

func (a *Application) index(router.Query) Renderable {
	routes := a.router.Routes()
	return Div(ForEach(len(routes), func(i int) Renderable {
		return P(A(Href("#"+routes[i].Path), Text(routes[i].Path)))
	}))
}

func (a *Application) Run() {
	a.router.Start()
	select {}
}
