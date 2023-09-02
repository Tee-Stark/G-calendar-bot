package services

import (
	"context"
	"encoding/json"
	"g_calendar_pal/bot/config"
	"g_calendar_pal/bot/utils"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var oauthConfig = config.InitOauth()

func GenerateAuthLink(w http.ResponseWriter, tgUsername string) string {
	oauthState := generateState(w, tgUsername)

	authUrl := oauthConfig.AuthCodeURL(oauthState)
	return authUrl
}

func generateState(w http.ResponseWriter, tgUsername string) string {
	expiry := time.Now().Add(30 * 24 * time.Hour)

	state := utils.GenerateShortID() + tgUsername
	cookie := http.Cookie{
		Name:    "OauthState",
		Value:   state,
		Expires: expiry,
	}
	// log.Println(cookie)
	// http.Delete
	http.SetCookie(w, &cookie)

	return state
}

func GetGoogleTokens(ctx context.Context, authCode string) (*oauth2.Token, error) {
	tokens, err := oauthConfig.Exchange(ctx, authCode)
	if err != nil {
		return &oauth2.Token{}, err
	}

	return tokens, nil
}

func GetUserEmailFromProfile(ctx context.Context, userTokens *oauth2.Token) (string, error) {
	client := oauthConfig.Client(ctx, userTokens)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Error while making request to client")
		return "", err
	}

	defer resp.Body.Close()

	var profile utils.GoogleProfile
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body")
		return "", err
	}

	log.Printf("Body: %s", string(body))
	err = json.Unmarshal(body, &profile)
	if err != nil {
		log.Println("Error while unmarshalling JSON")
		return "", err
	}
	log.Printf("Profile email: %v", profile.Email)

	return profile.Email, nil
}
