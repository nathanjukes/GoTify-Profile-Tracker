package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	// Attaching middlewares
	chiDispatcher.Use(middleware.RequestID)
	chiDispatcher.Use(middleware.RealIP)
	chiDispatcher.Use(middleware.Logger)
	chiDispatcher.Use(middleware.Recoverer)

	return &chiRouter{}
}

func (*chiRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, chiDispatcher))
}
