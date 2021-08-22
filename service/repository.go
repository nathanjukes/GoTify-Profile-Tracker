package service

import (
	"gotify-profile-tracker/entity"

	"gorm.io/gorm"
)

type Repository interface {
	DatabaseRefresher()
	updateCurrentlyPlaying(*gorm.DB)
	GetUserActivity(id string) []entity.Activity
	GetActivity() []entity.Activity
}
