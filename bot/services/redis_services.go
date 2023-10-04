package services

import (
	"context"
	"g_calendar_pal/bot/config"
	"g_calendar_pal/bot/utils"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/oauth2"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService() *RedisService {
	return &RedisService{
		client: config.InitRedis(),
	}
}

func (r *RedisService) SaveUserAuthTokens(email string, authTokens *oauth2.Token) {
	tokens := utils.SerializeUserTokenData(authTokens)

	saved := r.client.Set(context.Background(), email, tokens, 0)
	if saved.Err() != nil {
		log.Printf("Error saving user auth tokens: %v", saved.Err().Error())
	}
}

func (r *RedisService) GetUserAuthTokens(email string) *oauth2.Token {
	tokens, err := r.client.Get(context.Background(), email).Result()
	if err != nil {
		log.Printf("Error getting user authentication tokens: %v", tokens)
		return &oauth2.Token{}
	}
	userTokens := utils.DeserializeUserTokenData(tokens)
	return userTokens
}

func (r *RedisService) SaveUserState(userName string, userState utils.UserStateData) {
	state := utils.SerializeStateData(userState)

	saved := r.client.Set(context.Background(), userName, state, 0)
	if saved.Err() != nil {
		log.Printf("Error saving user state: %v", saved.Err())
	}
}

func (r *RedisService) GetUserState(userName string) utils.UserStateData {
	state, err := r.client.Get(context.Background(), userName).Result()
	if err != nil {
		log.Printf("Error getting user state: %v", err)
		return utils.UserStateData{}
	}
	stateData := utils.DeserializeStateData(state)

	return stateData
}

func (r *RedisService) SaveSession(sessionId string, event *utils.EventData) error {
	eventData := utils.SerializeEventData(event)
	saved := r.client.Set(context.Background(), sessionId, eventData, 0)
	if saved.Err() != nil {
		log.Printf("Error saving session: %v", saved.Err())
		return saved.Err()
	}
	log.Println("Session saved, ID: ", sessionId, "Event saved: ", eventData)
	return nil
}

func (r *RedisService) GetSession(sessionId string) *utils.EventData {
	event, err := r.client.Get(context.Background(), sessionId).Result()

	if err != nil {
		log.Printf("Error getting session: %v", err)
		return nil
	}

	eventData := utils.DeserializeEventData(event)

	return eventData
}
