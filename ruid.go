package ruid

import (
	"net/http"

	"github.com/egoholic/router"

	"github.com/egoholic/router/params"
	"github.com/egoholic/ruid/uid"
)

type handler struct {
	prefixes []string
	original http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("X-RUID", uid.New(h.prefixes))
	h.original.ServeHTTP(w, r)
}

func WrapHandlerWithRUID(prefixes []string, original http.Handler) http.Handler {
	return &handler{prefixes, original}
}

func WrapHandlerFnWithRUID(prefixes []string, original http.HandlerFunc) router.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *params.Params) {
		r.Header.Add("X-RUID", uid.New(prefixes))
		original(w, r)
	}
}
