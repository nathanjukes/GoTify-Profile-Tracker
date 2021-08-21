package service

import (
	"fmt"
	"gotify-profile-tracker/utility"
	"log"
	"time"

	gotify "github.com/nathanjukes/GoTify-BuddyList/buddyList"
	"gorm.io/gorm"
)

type repository struct{}

// Dependency received here, SpotifyRequester service returned
func NewSpotifyRepository() Repository {
	return &repository{}
}

//
// Keeps Database updated
//
func (r *repository) DatabaseRefresher(db *gorm.DB) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Database update imminent.")
				r.updateCurrentlyPlaying(db)
				log.Println("Database update complete.")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (*repository) updateCurrentlyPlaying(db *gorm.DB) {
	// Get data
	cookie := "AQC75BrTMdaqZSerrutLz7Bv8iLjIN_7MAc_QotbSU0hH9w5IQcBTN37Zpc274TVKYvol4r0W14ZCHk1atln5Q-WtsMJp3p28CfeE9RblQ8vJ-GX6MabsRxsud_d0jZe_SP7g2Vl5GqkqHrzwc67CpLsnr3MET_T_Y2_5134"

	gotify.NewScopedInstance(cookie)

	g := gotify.NewScopedInstance(cookie)
	bl, _ := g.BuddyList()
	friends := bl.Friends

	for _, i := range friends {
		activity := utility.GetActivity(i)
		fmt.Println(activity.Track.Album.Name)
		db.Save(&activity)
	}
}
