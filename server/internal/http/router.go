package http

import (
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// newRouter creates a router and connects the endpoints with the given server and its methods.
func (s *Server) newRouter(fileServerDir string) *httprouter.Router {
	logMe := func(p string) string {
		s.logger.Print(ecs.Msg("registered endpoint"), log.V("url.path", p))
		return p
	}

	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, logMe("/blub"), func(writer http.ResponseWriter, request *http.Request) {
		s.logger.Print(ecs.Msg("hello world"))
	})
	router.HandlerFunc(http.MethodGet, logMe("/api/v1/poll/version"), s.pollVersion)

	if fileServerDir != "" {
		router.NotFound = http.FileServer(http.Dir(logMe(fileServerDir)))
	}

	return router
}
