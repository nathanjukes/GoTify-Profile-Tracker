package main

import (
	"gotify-profile-tracker/controller"
	"gotify-profile-tracker/database"
	"gotify-profile-tracker/http"
	"gotify-profile-tracker/service"
	"log"
	"os"
)

func init() {
	// Environment variable setup
	os.Setenv("version", "development")

	// Personal
	os.Setenv("CLIENT_ID", "")
	os.Setenv("CLIENT_SECRET", "")
	os.Setenv("ACCESS_TOKEN", "")
	os.Setenv("REFRESH_TOKEN", "")
	os.Setenv("SDPC_COOKIE", "")
}

func main() {
	log.Println("start")

	r := http.NewChiRouter()

	// Setup DB
	Idb := database.GetPostgresDB()
	db := Idb.GetDB()

	// Keeps database updated
	repo := service.NewSpotifyRepository(db)
	repo.DatabaseRefresher()

	// Get Spotify service
	service := service.NewSpotifyService()

	// Initialise controllers
	rc := controller.NewRepoController(repo, service)
	uc := controller.NewUtilityController()
	ec := controller.NewErrorController()

	mc := controller.NewController(rc, uc, ec)

	// Starts serving
	serve(r, mc)
}

func serve(r http.Router, c controller.Controller) {
	r.Get("/", c.Utility.Health)
	r.Get("/health", c.Utility.Health)
	r.Get("/version", c.Utility.Version)

	r.Get("/user/{userID}", c.Repo.User)
	r.Get("/activity", c.Repo.Activity)

	r.Post("/user/{userID}", c.Repo.UserAdd)

	r.Serve(":8081")
}
