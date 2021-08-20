package main

import (
	"fmt"
	"gotify-profile-tracker/controller"
	"gotify-profile-tracker/http"
	"os"

	gotify "github.com/nathanjukes/GoTify-BuddyList/buddyList"
)

func init() {
	// Environment variable setup
	// TODO: Change this hardcoded implementation into a docker v.c implementation
	os.Setenv("version", "development")
}

func main() {
	fmt.Println("start")

	cookie := "AQC75BrTMdaqZSerrutLz7Bv8iLjIN_7MAc_QotbSU0hH9w5IQcBTN37Zpc274TVKYvol4r0W14ZCHk1atln5Q-WtsMJp3p28CfeE9RblQ8vJ-GX6MabsRxsud_d0jZe_SP7g2Vl5GqkqHrzwc67CpLsnr3MET_T_Y2_5134"

	gotify.NewScopedInstance(cookie)

	g := gotify.NewScopedInstance(cookie)
	bl, _ := g.BuddyList()
	for _, i := range bl.Friends {
		fmt.Println(i.User.Name + i.Track.Name)
	}

	r := http.NewChiRouter()
	c := controller.NewUtilityController()
	serve(r, c)
}

func serve(r http.Router, c controller.UtilityController) {
	r.Get("/health", c.Health)
	r.Get("/version", c.Version)

	r.Serve(":8080")
}
