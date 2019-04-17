package healthcheck

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Checker interface {
	Ready() bool
	Live() bool
}

func New(c Checker) *httprouter.Router {
	router := httprouter.New()
	router.GET("/_health", ready(c))
	router.GET("/_ready", live(c))
	return router
}

func ready(c Checker) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if !c.Ready() {
			w.WriteHeader(500) // Non 200
			return
		}
		w.WriteHeader(200)
	}
}

func live(c Checker) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if !c.Live() {
			w.WriteHeader(500) // Non 200
			return
		}
		w.WriteHeader(200)
	}
}