package utils

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type EventData struct {
	EventName        string   `json:"event_name"`
	EventDescription string   `json:"event_description"`
	EventStartDate   string   `json:"event_start_date"`
	EventEndDate     string   `json:"event_end_date"`
	EventStartTime   string   `json:"event_start_time"`
	EventEndTime     string   `json:"event_end_time"`
	EventTimeZone    string   `json:"event_time_zone"`
	EventLocation    string   `json:"event_location"`
	EventAttendees   []string `json:"event_attendees"`
}

func NewEvent() *EventData {
	return &EventData{
		EventAttendees: make([]string, 0),
	}
}

func SerializeEventData(eventData *EventData) string {
	e, err := json.Marshal(&eventData)
	if err != nil {
		log.Printf("Error serializing event data: %v", err)
		return ""
	}

	return string(e)
}

func DeserializeEventData(eventData string) *EventData {
	var e EventData
	err := json.Unmarshal([]byte(eventData), &e)
	if err != nil {
		log.Printf("Error deserializing event data: %v", err)
	}

	return &e
}

func GenerateShortID() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.New(rand.NewSource(time.Now().UnixNano()))

	sessionID := make([]byte, 8)
	for i := range sessionID {
		sessionID[i] = chars[rand.Intn(len(chars))]
	}

	return string(sessionID)
}
