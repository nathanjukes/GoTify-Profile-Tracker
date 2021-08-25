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
	os.Setenv("CLIENT_ID", "6a90e4b03ccb4c63b7784e5e30d89eb9")
	os.Setenv("CLIENT_SECRET", "2b93e5e1ae1c45259694f4d5a3f914f5")
	os.Setenv("ACCESS_TOKEN", "BQBN7H_Mlvx1PXYvXkrkLqarXo36LH3qeKX_WoRTwQG-gzs3qHUFHhAbOTFRgP15TX7fnokLRP4C5s4nZDnviUcaGD5Ee6P9rpj0rpMK763daieVRZuhX7o1Jhi8ZZ-PKQj8o2FFFAHm9NFeqgVg")
	os.Setenv("REFRESH_TOKEN", "AQBD80DsPjdEu3FaauvdtbkhlMwATuplxXRf4Ej_OCFhg4k22R0rJ7lYTTXDn9SOAMLLD35otyYTznPY0Fj17Mv4cLVAJFCv6TWjYyMewkXcCn_k4pVcoGk5fzxmnKBW4es")
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
