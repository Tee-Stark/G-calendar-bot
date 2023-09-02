package utils

import (
	"encoding/json"
	"log"
)

var UserStates = map[string]string{
	"start":        "start",
	"auth":         "auth",
	"authSuccess":  "authSuccess",
	"authError":    "authError",
	"createEvent":  "0",
	"createEvent1": "1",
	"createEvent2": "2",
	"createEvent3": "3",
	"createEvent4": "4",
	"createEvent5": "5",
	"createEvent6": "6",
	"createEvent7": "7",
	"createEvent8": "8",
	"createEvent9": "9",
	"eventCreated": "eventCreated",
}

type UserStateData struct {
	UserEmail string `json:"user_email"`
	State     string `json:"state"`
	SessionID string `json:"session_id"`
}

func SerializeStateData(stateData UserStateData) string {
	s, err := json.Marshal(stateData)
	if err != nil {
		log.Printf("Error serializing state data: %v", err)
		return ""
	}

	return string(s)
}

func DeserializeStateData(stateData string) UserStateData {
	var s UserStateData
	err := json.Unmarshal([]byte(stateData), &s)
	if err != nil {
		log.Printf("Error deserializing state data: %v", err)
	}

	return s
}
