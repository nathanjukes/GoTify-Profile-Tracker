package main

import (
	"fmt"
	"gotify-profile-tracker/controller"
	"gotify-profile-tracker/database"
	"gotify-profile-tracker/http"
	"gotify-profile-tracker/service"
	"os"
)

func init() {
	// Environment variable setup
	// TODO: Change this hardcoded implementation into a docker v.c implementation
	os.Setenv("version", "development")
}

func main() {
	fmt.Println("start")

	r := http.NewChiRouter()

	// Setup DB
	Idb := database.GetPostgresDB()
	db := Idb.GetDB()

	// Keeps database updated
	repo := service.NewSpotifyRepository(db)
	repo.DatabaseRefresher()

	// Initialise controllers
	rc := controller.NewRepoController(repo)
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

	r.Serve(":8080")
}
