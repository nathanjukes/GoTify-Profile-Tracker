package controller

import (
	"encoding/json"
	"net/http"
)

type errorController struct{}

func NewErrorController() ErrorController {
	return &errorController{}
}

//
// Error page
//

type errorPage struct {
	Path   string
	Status int
}

func (*errorController) ErrorNotFound(w http.ResponseWriter, r *http.Request) {
	errId := r.URL
	e := errorPage{errId.Path, http.StatusNotFound}
	json.NewEncoder(w).Encode(e)
}
