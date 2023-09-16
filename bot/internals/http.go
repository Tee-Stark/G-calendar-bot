package internals

import (
	"g_calendar_pal/bot/services"
	"g_calendar_pal/bot/utils"
	"g_calendar_pal/views"
	"log"
	"net/http"
)

func OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("username")
	loginUrl := services.GenerateAuthLink(w, userName)

	http.Redirect(w, r, loginUrl, http.StatusTemporaryRedirect)
}

func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")

	// verify state
	state := r.FormValue("state")
	cookie, err := r.Cookie("OauthState")
	if err != nil {
		log.Printf("An error occured while getting cookie: %v", err)
		return
	}

	// log.Printf("State: %s, Cookie.Value: %v", state, cookie)

	if state != cookie.Value {
		log.Println("Invalid state")
		http.Redirect(w, r, "/auth-error", http.StatusTemporaryRedirect)
		return
	}
	// store tokens and user email in Redis here
	userTokens, err := services.GetGoogleTokens(r.Context(), authCode)
	if err != nil {
		log.Printf("An error occured while getting tokens: %v", err)
	}
	// Get user email from people API
	userEmail, err := services.GetUserEmailFromProfile(r.Context(), userTokens)
	if err != nil {
		log.Printf("An error occured while getting user email: %v", err)
	}
	// store token and state in Redis here
	services.SaveUserAuthTokens(userEmail, userTokens)
	// update user state
	authSuccessState := utils.UserStateData{
		UserEmail: userEmail,
		State:     utils.UserStates["authSuccess"],
	}
	services.SaveUserState(state[8:], authSuccessState)
	log.Println("Authentication successful")

	http.Redirect(w, r, "/auth-success", http.StatusTemporaryRedirect)
}

func AuthSuccessHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Authentication successful"))
	views.RenderAuthResponse(w, "auth.html", "success")
}

func AuthErrorHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Authentication failed"))
	views.RenderAuthResponse(w, "auth.html", "")
}
