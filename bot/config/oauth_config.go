package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitOauth() *oauth2.Config {
	LoadEnv()

	CLIENT_ID := os.Getenv("G_CLIENT_ID")
	CLIENT_SECRET := os.Getenv("G_CLIENT_SECRET")
	REDIRECT_URL := os.Getenv("G_REDIRECT_URL")

	oauthConfig := oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL:  REDIRECT_URL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar.events",
			"https://www.googleapis.com/auth/calendar.events.owned",
		},
		Endpoint: google.Endpoint,
	}
	return &oauthConfig
}
