package http

import (
	"context"
	"fmt"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"net/http"
	"time"
)

// Server is the rest service.
type Server struct {
	host    string
	port    int
	httpSrv *http.Server
	dir     string
	logger  log.Logger
}

// NewServer prepares a new Server instance.
func NewServer(logger log.Logger, host string, port int, dir string) *Server {
	s := &Server{
		host:   host,
		port:   port,
		logger: logger,
		dir:    dir,
	}

	return s
}

// Run launches the server
func (s *Server) Run() error {
	router := s.newRouter(s.dir)

	s.httpSrv = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.host, s.port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      router,
	}

	s.logger.Print(ecs.Msg("starting"), ecs.ServerAddress(s.host), ecs.ServerPort(s.port))
	err := s.httpSrv.ListenAndServe()

	if err == http.ErrServerClosed {
		s.logger.Print(ecs.Msg("stopped"))
		return nil
	}

	return err
}

// Stop signals the server to halt gracefully.
func (s *Server) Stop() {
	if err := s.httpSrv.Shutdown(context.Background()); err != nil {
		s.logger.Print(ecs.Msg("failed to shutdown"), ecs.ErrMsg(err))
	}
}
