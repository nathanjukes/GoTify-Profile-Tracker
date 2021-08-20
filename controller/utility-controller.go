package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type utilityController struct{}

func NewUtilityController() UtilityController {
	return &utilityController{}
}

//
// Utilities
//

//
// GET /health
//
func (*utilityController) Health(w http.ResponseWriter, r *http.Request) {
	type health struct {
		Type    string    `json:"type"`
		Message string    `json:"message"`
		Healthy bool      `json:"healthy"`
		Time    time.Time `json:"time"`
	}
	h := health{"Backend Health Check", "Healthy", true, time.Now()}
	json.NewEncoder(w).Encode(h)
}

//
//  GET /version
//
func (u *utilityController) Version(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Version string    `json:"version"`
		Time    time.Time `json:"time"`
	}{
		os.Getenv("version"), // Environment variable
		time.Now(),
	})
}
