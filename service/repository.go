package service

import "gorm.io/gorm"

type Repository interface {
	DatabaseRefresher(*gorm.DB)
	updateCurrentlyPlaying(*gorm.DB)
}
