package utility

import (
	"gotify-profile-tracker/entity"

	types "github.com/nathanjukes/GoTify-BuddyList/entity"
)

func GetActivity(f types.Friends) entity.Activity {
	album := entity.Album{
		AlbumURI: f.Track.Album.URI,
		Name:     f.Track.Album.Name,
	}

	artist := entity.Artist{
		ArtistURI: f.Track.Artist.URI,
		Name:      f.Track.Artist.Name,
	}

	track := entity.Track{
		TrackURI: f.Track.URI,
		Name:     f.Track.Name,
		ImageURL: f.Track.ImageURL,
		Album:    album,
		Artist:   artist,
	}

	user := entity.User{
		UserURI:  f.User.URI,
		Name:     f.User.Name,
		ImageURL: f.User.ImageURL,
	}

	a := entity.Activity{
		Timestamp: f.Timestamp,
		Time:      f.Time,
		User:      user,
		Track:     track,
	}

	return a
}
