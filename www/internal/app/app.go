package app

import (
	"github.com/golangee/forms-example/www/internal/build"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

type Application struct {
	logger log.Logger
}

func NewApplication() *Application {
	return &Application{
		logger: log.NewLogger(ecs.Log("application")),
	}
}

func (a *Application) Run(){
	a.logger.Print(ecs.Msg("application is running14"), log.V("build.commit", build.Commit))
}
