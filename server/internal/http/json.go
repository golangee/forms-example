package http

import (
	"encoding/json"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"net/http"
	"reflect"
)

// writeJson is a helper and just tries to serialize the response as json.
func writeJson(w http.ResponseWriter, r *http.Request, obj interface{}) {
	buf, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.FromContext(r.Context()).Print(ecs.Msg("failed to marshal json response"), ecs.ErrMsg(err), log.V("type", reflect.TypeOf(obj).String()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(buf); err != nil {
		log.FromContext(r.Context()).Print(ecs.Msg("failed to write Json response"), ecs.ErrMsg(err))
	}
}
