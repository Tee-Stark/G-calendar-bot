package mocks

import (
	"g_calendar_pal/bot/utils"

	"golang.org/x/oauth2"
)

type RedisMock struct {
	client string
}

func NewRedisMock() *RedisMock {
	return &RedisMock{
		client: "yes",
	}
}

func (r *RedisMock) SaveUserAuthTokens(email string, authTokens *oauth2.Token) {

}

func (r *RedisMock) GetUserAuthTokens(email string) *oauth2.Token {

}

func (r *RedisMock) SaveUserState(userName string, userState utils.UserStateData) {

}

func (r *RedisMock) GetUserState(userName string) utils.UserStateData {

}

func (r *RedisMock) SaveSession(sessionId string, event *utils.EventData) error {

}

func (r *RedisMock) GetSession(sessionId string) *utils.EventData {

}

type RedisService interface {
	SaveUserAuthTokens(email string, authTokens *oauth2.Token)
	GetUserAuthTokens(email string) *oauth2.Token
	SaveUserState(userName string, userState utils.UserStateData)
	GetUserState(userName string) utils.UserStateData
	SaveSession(sessionId string, event *utils.EventData) error
	GetSession(sessionId string) *utils.EventData
}
