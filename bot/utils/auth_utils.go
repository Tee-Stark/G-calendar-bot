package utils

import (
	"encoding/json"
	"log"

	"golang.org/x/oauth2"
)

type GoogleProfile struct {
	Email string `json:"email"`
}

func SerializeUserTokenData(userToken *oauth2.Token) string {
	s, err := json.Marshal(&userToken)
	if err != nil {
		log.Printf("Error serializing user tokens: %v", err)
		return ""
	}

	return string(s)
}

func DeserializeUserTokenData(userToken string) *oauth2.Token {
	var t *oauth2.Token
	err := json.Unmarshal([]byte(userToken), &t)
	if err != nil {
		log.Printf("Error deserializing user tokens: %v", err)
	}

	return t
}
