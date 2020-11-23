package app

import (
	"errors"
	"fmt"
	builder2 "github.com/golangee/forms-example/server/internal/builder"
	"github.com/golangee/forms-example/server/internal/http"
	"github.com/golangee/forms-example/server/internal/livebuilder"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

type Application struct {
	server  *http.Server
	logger  log.Logger
	builder *livebuilder.Builder
	tmpDir  string
}

func NewApplication(host string, port int, wwwDir string) (*Application, error) {
	tmpDir := filepath.Join(os.TempDir(), "golangee-forms-livebuilder")
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return nil, err
	}

	a := &Application{}
	a.initCloseListener()
	a.logger = log.NewLogger(ecs.Log("application"))

	a.logger.Print(ecs.Msg("build dir " + tmpDir))
	wwwBuildDir := filepath.Join(tmpDir, "www")

	a.server = http.NewServer(log.WithFields(a.logger, ecs.Log("httpserver")), host, port, wwwBuildDir)
	builder, err := livebuilder.NewBuilder(wwwBuildDir, wwwDir)
	if err != nil {
		return nil, err
	}
	a.builder = builder
	if err := a.builder.Build(); err != nil {
		buildErr := builder2.BuildErr{}
		if errors.As(err, &buildErr) {
			a.logger.Print(ecs.ErrMsg(err))
		} else {
			return nil, fmt.Errorf("unable to create initial build: %w", err)
		}
	}

	return a, nil
}

func (a *Application) initCloseListener() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		a.server.Stop()
	}()
}

func (a *Application) Run() error {
	defer func() {
		a.logger.Print(ecs.Msg("exiting"))
	}()

	return a.server.Run()
}

func (a *Application) Close() error {
	a.server.Stop()
	return os.RemoveAll(a.tmpDir)
}
