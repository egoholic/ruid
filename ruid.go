package ruid

import (
	"net/http"

	"github.com/egoholic/ruid/uid"
)

type handler struct {
	label    string
	original http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("RUID", uid.New(h.label))
	h.original.ServeHTTP(w, r)
}

func WrapHandlerWithRUID(label string, original http.Handler) http.Handler {
	return &handler{label, original}
}

func WrapHandlerFnWithRUID(label string, original http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("RUID", uid.New(label))
		original(w, r)
	}
}
