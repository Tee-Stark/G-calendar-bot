package services

import (
	"bytes"
	"encoding/json"
	"g_calendar_pal/bot/config"
	"g_calendar_pal/bot/utils"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"google.golang.org/api/calendar/v3"
)

func CreateCalendarEvent(userEmail string, event utils.EventData) error {
	// get the user's access token from redis using userEmail
	accessToken := GetUserAuthTokens(userEmail)
	// extract attendees
	var eventAttendees []*calendar.EventAttendee
	// append current user and make him organizer
	eventAttendees = append(eventAttendees, &calendar.EventAttendee{Email: userEmail, Organizer: true})
	for _, attendee := range event.EventAttendees {
		eventAttendee := &calendar.EventAttendee{
			Email: attendee,
		}
		eventAttendees = append(eventAttendees, eventAttendee)
	}

	log.Println(&eventAttendees)

	// eventID := uuid.New().String()
	requestID := uuid.New().String()
	startDateTime, _ := utils.GenerateDateTime(event.EventStartDate, event.EventStartTime)
	endDateTime, _ := utils.GenerateDateTime(event.EventEndDate, event.EventEndTime)

	Calendar := config.InitCalendar(accessToken)

	newGoogleEvent := Calendar.Events.Insert("primary", &calendar.Event{
		Summary: event.EventName,
		Start: &calendar.EventDateTime{
			DateTime: startDateTime,
			TimeZone: "Africa/Lagos", //event.EventTimeZone,
		},
		End: &calendar.EventDateTime{
			DateTime: endDateTime,
			TimeZone: "Africa/Lagos", //event.EventTimeZone,
		},
		Description: event.EventDescription,
		Location:    event.EventLocation,
		Attendees:   eventAttendees,
		// meeting link
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: requestID,
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		},
		Reminders: &calendar.EventReminders{
			UseDefault: true,
		},
	})

	// set user token, and sent updates to all attendees
	_, err := newGoogleEvent.ConferenceDataVersion(1).SendUpdates("all").Do()
	if err != nil {
		log.Printf("Error creating event: %v", err.Error())
		return err
	}

	log.Println("Event created successfully")
	return nil
}

// alternative implementation to directly call API using HTTP
func CreateCalendarEventHTTP(userEmail string, event utils.EventData) error {
	accessToken := GetUserAuthTokens(userEmail)

	var eventAttendees = make([]map[string]interface{}, 0)
	// append current user and make him organizer
	organizer := map[string]interface{}{
		"email":     userEmail,
		"organizer": true,
	}

	eventAttendees = append(eventAttendees, organizer)
	for _, attendee := range event.EventAttendees {
		att := map[string]interface{}{
			"email": attendee,
		}
		eventAttendees = append(eventAttendees, att)
	}

	requestID := uuid.New().String()
	// event data to submit
	startDateTime, _ := utils.GenerateDateTime(event.EventStartDate, event.EventStartTime)
	endDateTime, _ := utils.GenerateDateTime(event.EventEndDate, event.EventEndTime)
	eventData := map[string]interface{}{
		"summary": event.EventName,
		"start": map[string]interface{}{
			"dateTime": startDateTime,
			"timeZone": "Africa/Lagos",
		},
		"end": map[string]interface{}{
			"dateTime": endDateTime,
			"timeZone": "Africa/Lagos",
		},
		"description": event.EventDescription,
		"location":    event.EventLocation,
		"attendees":   eventAttendees,
		"conferenceData": map[string]interface{}{
			"createRequest": map[string]interface{}{
				"conferenceSolutionKey": map[string]interface{}{
					"type": "hangoutsMeet",
				},
				"requestId": requestID,
			},
		},
		"reminders": map[string]interface{}{
			"useDefault": true,
		},
	}

	eventDataJSON, err := json.Marshal(eventData)
	if err != nil {
		log.Printf("Error marshalling event data: %v", err)
		return err
	}
	// create HTTP request
	url := "https://www.googleapis.com/calendar/v3/calendars/primary/events"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(eventDataJSON))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
	}

	accessToken.SetAuthHeader(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("HTTP request failed with status code: %v", resp.StatusCode)
		log.Printf("Response body: %v", string(respBody))
		return err
	}

	log.Println(string(respBody))
	return nil
}
