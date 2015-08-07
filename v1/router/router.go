package router

import (
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/dnl-jst/goard-daemon/v1/routeHandler"
	"go.iondynamics.net/webapp"
)

func New() *mux.Router {
	return provision(mux.NewRouter().StrictSlash(true))
}

func provision(r *mux.Router) *mux.Router {

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/v1/", http.StatusFound)
	})

	v1Post := r.PathPrefix("/v1").Methods("POST").Subrouter()
	v1Post.Handle("/check-http", webapp.Handler(handler.CheckHttp))

	return r
}
