package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Timestamp int64     `json:"timestamp"`
	Time      time.Time `json:"time"`
	UserID    int
	User      User `json:"user"`
	TrackID   int
	Track     Track `json:"track"`
}

type User struct {
	gorm.Model
	UserURI  string `json:"uri"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

type Album struct {
	gorm.Model
	AlbumURI string `json:"uri"`
	Name     string `json:"name"`
}

type Artist struct {
	gorm.Model
	ArtistURI string `json:"uri"`
	Name      string `json:"name"`
}

type Track struct {
	gorm.Model
	TrackURI string `json:"uri"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	AlbumID  int
	Album    Album `json:"album"`
	ArtistID int
	Artist   Artist `json:"artist"`
}
