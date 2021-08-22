package controller

import "net/http"

type RepoController interface {
	User(w http.ResponseWriter, r *http.Request)
	Activity(w http.ResponseWriter, r *http.Request)
}

type UtilityController interface {
	Health(w http.ResponseWriter, r *http.Request)
	Version(w http.ResponseWriter, r *http.Request)
}

type ErrorController interface {
	ErrorNotFound(w http.ResponseWriter, r *http.Request)
}

//
// Main Controller to link Repo, Utility & Error controllers with DI
//

type Controller struct {
	Repo    RepoController
	Utility UtilityController
	Error   ErrorController
}

func NewController(r RepoController, u UtilityController, e ErrorController) Controller {
	return Controller{r, u, e}
}
