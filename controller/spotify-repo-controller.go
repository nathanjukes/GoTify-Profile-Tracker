package controller

import (
	"encoding/json"
	"gotify-profile-tracker/service"
	"net/http"

	"github.com/go-chi/chi"
)

type repoController struct {
	repo    service.Repository
	service service.Service
}

func NewRepoController(r service.Repository, s service.Service) RepoController {
	return &repoController{repo: r, service: s}
}

//
// Data related methods
//

func (s *repoController) User(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(s.repo.GetUserActivity(userID))
}

func (s *repoController) UserAdd(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	uniqueID := userID[13:]
	s.service.FollowUser(uniqueID)
	json.NewEncoder(w).Encode(s.repo.GetUserActivity(userID))
}

func (s *repoController) Activity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(s.repo.GetActivity())
}
