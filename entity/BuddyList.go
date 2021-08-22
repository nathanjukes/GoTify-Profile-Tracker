package entity

import (
	"time"
)

type Activity struct {
	Timestamp int64     `json:"timestamp" gorm:"primaryKey;autoIncrement:false"`
	Time      time.Time `json:"time"`
	UserID    string    `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	User      User      `json:"user" gorm:"foreignkey:UserID"`
	TrackID   string    `json:"track_id"`
	Track     Track     `json:"track" gorm:"foreignkey:TrackID"`
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
	AlbumID  string `json:"album_id"`
	Album    Album  `json:"album" gorm:"foreignkey:AlbumID"`
	ArtistID string `json:"artist_id"`
	Artist   Artist `json:"artist" gorm:"foreignkey:ArtistID"`
}
