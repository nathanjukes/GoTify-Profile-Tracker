package controller

import (
	"encoding/json"
	"gotify-profile-tracker/service"
	"net/http"

	"github.com/go-chi/chi"
)

type repoController struct {
	repo service.Repository
}

func NewRepoController(r service.Repository) RepoController {
	return &repoController{repo: r}
}

//
// Data related methods
//

func (s *repoController) User(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	json.NewEncoder(w).Encode(s.repo.GetUserActivity(userID))
}

func (s *repoController) Activity(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.repo.GetActivity())
}
