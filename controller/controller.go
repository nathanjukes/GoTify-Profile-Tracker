package controller

import "net/http"

type UtilityController interface {
	Health(w http.ResponseWriter, r *http.Request)
	Version(w http.ResponseWriter, r *http.Request)
}
