package entity

import (
	"time"
)

type Activity struct {
	Timestamp int64     `json:"timestamp" gorm:"primaryKey;autoIncrement:false"`
	Time      time.Time `json:"time"`
	UserID    string    `gorm:"primaryKey;autoIncrement:false"`
	User      User      `json:"user"`
	TrackID   string
	Track     Track `json:"track"`
}

type User struct {
	UserURI  string `json:"uri" gorm:"primary_key"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

type Album struct {
	AlbumURI string `json:"uri" gorm:"primary_key"`
	Name     string `json:"name"`
}

type Artist struct {
	ArtistURI string `json:"uri" gorm:"primary_key"`
	Name      string `json:"name"`
}

type Track struct {
	TrackURI string `json:"uri" gorm:"primary_key"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	AlbumID  string
	Album    Album `json:"album"`
	ArtistID string
	Artist   Artist `json:"artist"`
}
