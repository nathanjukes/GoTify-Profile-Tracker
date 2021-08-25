package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gotify-profile-tracker/entity"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type service struct {
	a entity.Auth
}

// Dependency received here, SpotifyRequester service returned
func NewSpotifyService() Service {
	return &service{a: entity.Auth{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		Scope:        "user-follow-modify",
	}}
}

func (s *service) FollowUser(id string) error {
	err := refreshToken(&s.a)

	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("https://api.spotify.com/v1/me/following?type=user&ids=%s", id)
	req, err := http.NewRequest(http.MethodPut, endpoint, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+s.a.AccessToken)

	// Get response from POST above
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.Status != "204" || err != nil {
		return fmt.Errorf("follower not added")
	} else {
		return nil
	}
}

func refreshToken(a *entity.Auth) error {
	endpoint := "https://accounts.spotify.com/api/token"

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(url.Values{"grant_type": {"refresh_token"}, "refresh_token": {a.RefreshToken}}.Encode()))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+getBase64EncodedAuth())

	// Get response from POST above
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	dat, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(dat, a)
	return nil
}

func getBase64EncodedAuth() string {
	return base64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID") + ":" + os.Getenv("CLIENT_SECRET")))
}
