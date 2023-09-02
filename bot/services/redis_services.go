package services

import (
	"context"
	"g_calendar_pal/bot/config"
	"g_calendar_pal/bot/utils"
	"log"

	"golang.org/x/oauth2"
)

var Redis = config.InitRedis()

func SaveUserAuthTokens(email string, authTokens *oauth2.Token) {
	tokens := utils.SerializeUserTokenData(authTokens)

	saved := Redis.Set(context.Background(), email, tokens, 0)
	if saved.Err() != nil {
		log.Printf("Error saving user auth tokens: %v", saved.Err().Error())
	}
}

func GetUserAuthTokens(email string) *oauth2.Token {
	tokens, err := Redis.Get(context.Background(), email).Result()
	if err != nil {
		log.Printf("Error getting user authentication tokens: %v", tokens)
		return &oauth2.Token{}
	}
	userTokens := utils.DeserializeUserTokenData(tokens)
	// log.Println("User tokens: ", tokens)
	return userTokens
}

func SaveUserState(userName string, userState utils.UserStateData) {
	state := utils.SerializeStateData(userState)

	saved := Redis.Set(context.Background(), userName, state, 0)
	if saved.Err() != nil {
		log.Printf("Error saving user state: %v", saved.Err())
		// return saved.Err()
	}
	// log.Println("User state updated correctly")
}

func GetUserState(userName string) utils.UserStateData {
	state, err := Redis.Get(context.Background(), userName).Result()
	if err != nil {
		log.Printf("Error getting user state: %v", err)
		return utils.UserStateData{}
	}
	stateData := utils.DeserializeStateData(state)

	return stateData
}

func SaveSession(sessionId string, event *utils.EventData) error {
	eventData := utils.SerializeEventData(event)

	log.Println("Session to save: ", sessionId, "Event to save: ", eventData)
	saved := Redis.Set(context.Background(), sessionId, eventData, 0)
	if saved.Err() != nil {
		log.Printf("Error saving session: %v", saved.Err())
		return saved.Err()
	}
	log.Println("Session saved correctly")
	return nil
}

func GetSession(sessionId string) *utils.EventData {
	event, err := Redis.Get(context.Background(), sessionId).Result()

	log.Println("Get Session: ", event)
	if err != nil {
		log.Printf("Error getting session: %v", err)
		return nil
	}

	eventData := utils.DeserializeEventData(event)

	return eventData
}
