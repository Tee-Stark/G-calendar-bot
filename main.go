package main

import (
	"g_calendar_pal/bot/config"
	"g_calendar_pal/bot/internals"
	"log"
	"net/http"
)

func main() {
	// load .env
	config.LoadEnv()
	// start http server in different goroutine
	go startHTTPServer()

	// start bot and services
	log.Println("G-CALENDAR BOT IS RUNNING...")
	bot := config.InitBot()
	internals.HandleBot(bot)
}

func startHTTPServer() {
	http.HandleFunc("/oauth-google", internals.OauthGoogleLogin)
	http.HandleFunc("/oauth-redirect", internals.OAuthCallbackHandler)
	http.HandleFunc("/auth-success", internals.AuthSuccessHandler)
	http.HandleFunc("/auth-error", internals.AuthErrorHandler)

	// Start HTTP server
	if err := http.ListenAndServe(":4545", nil); err != nil {
		log.Fatal(err)
	}
}
