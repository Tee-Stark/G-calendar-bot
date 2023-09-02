package config

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func InitCalendar(token *oauth2.Token) *calendar.Service {
	LoadEnv()
	ctx := context.Background()
	// apiKey := os.Getenv("G_API_KEY")
	// credentialFile := os.Getenv("G_CREDENTIALS_PATH")
	tokenSource := oauth2.StaticTokenSource(token)

	cal, err := calendar.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatalf("An error occured while initializing Google Calendar: %v", err)
	}

	log.Println("Initialized Google Calendar")
	return cal
}
