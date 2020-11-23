package http

import (
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"net/http"
)

func (s *Server) pollVersion(w http.ResponseWriter, r *http.Request) {
	log.FromContext(r.Context()).Print(ecs.Msg("registered long poll"))

	c := s.await()
	select {
	case version := <-c:
		type Version struct {
			Version string
		}
		log.FromContext(r.Context()).Print(ecs.Msg("returning " + version))
		writeJson(w, r, Version{Version: version})
	}
}
