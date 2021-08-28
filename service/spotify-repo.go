package service

import (
	"gotify-profile-tracker/entity"
	"gotify-profile-tracker/utility"
	"log"
	"os"
	"time"

	gotify "github.com/nathanjukes/GoTify-BuddyList/buddyList"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Dependency received here, SpotifyRequester service returned
func NewSpotifyRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

//
// Keeps Database updated
//
func (r *repository) DatabaseRefresher() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Database update imminent.")
				r.updateCurrentlyPlaying(r.db)
				log.Println("Database update complete.")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (*repository) updateCurrentlyPlaying(db *gorm.DB) {
	// Get activity
	cookie := os.Getenv("SDPC_COOKIE")

	gotify.NewScopedInstance(cookie)
	g := gotify.NewScopedInstance(cookie)
	bl, _ := g.BuddyList()
	friends := bl.Friends

	// Insert activity into database
	for _, i := range friends {
		activity := utility.GetActivity(i)
		db.Save(&activity)
	}
}

func (r *repository) GetUserActivity(id string) []entity.Activity {
	var activity []entity.Activity

	err := r.db.Preload("User").Preload("Track").Preload("Track.Album").Preload("Track.Artist").Where("user_id = ?", id).Order("time desc").Find(&activity)
	if err.Error != nil {
		log.Fatalf(err.Error.Error())
	}

	return activity
}

func (r *repository) GetActivity() []entity.Activity {
	var activity []entity.Activity

	err := r.db.Preload("User").Preload("Track").Preload("Track.Album").Preload("Track.Artist").Order("time desc").Find(&activity)
	if err.Error != nil {
		log.Fatalf(err.Error.Error())
	}

	return activity
}
