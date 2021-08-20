package controller

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type repoController struct{}

func NewRepoController() RepoController {
	return &repoController{}
}

//
// Data related methods
//

func (s *repoController) User(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	log.Println(userID)
}
