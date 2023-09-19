package services

import (
	"fmt"
	"g_calendar_pal/bot/utils"
	"strconv"
	"strings"
	"text/template"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Start Command to Handle the /start command
func HandleStart(update tgbotapi.Update) tgbotapi.MessageConfig {
	responseTemp := utils.ResponseTemplates["start"]
	userName := update.Message.From.UserName

	// parse username and put it in the response
	tmpl := template.Must(template.New("response text").Parse(responseTemp))
	var responseText strings.Builder
	if err := tmpl.Execute(&responseText, map[string]string{"username": userName}); err != nil {
		log.Printf("An error occured while executing template: %v", err)
	}

	textMsg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText.String())
	textMsg.ReplyToMessageID = update.Message.MessageID
	googleSigninURL := "http://127.0.0.1:4545/oauth-google?username=" + userName

	keyMsg := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Google Signin", googleSigninURL)),
	)
	textMsg.ReplyMarkup = keyMsg

	SaveUserState(userName, utils.UserStateData{State: utils.UserStates["start"]})

	return textMsg
}

func HandleCreateEvent(update tgbotapi.Update, userStateData utils.UserStateData) tgbotapi.MessageConfig {
	var responseText string
	var textMsg tgbotapi.MessageConfig
	var sessionID string
	// var event *utils.EventData

	username := update.Message.From.UserName

	switch userStateData.State {
	case "0":
		// create new session and update userStateData
		sessionID = utils.GenerateShortID()
		event := utils.NewEvent()
		userStateData.SessionID = sessionID
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "1":
		// event name is the message sent
		eventName := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventName = eventName
		// save session and state again
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "2":
		// event description is the message sent
		eventDesc := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventDescription = eventDesc
		// save session and state again
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "3":
		// event start date is the message sent
		eventStartDate := update.Message.Text
		_, _, _, err := utils.ParseDate(eventStartDate)
		if err != nil {
			responseText = utils.ErrorResponses["dateError"]
			userStateData.State = utils.UserStates["createEvent3"]
			SaveUserState(username, userStateData)
			break
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventStartDate = eventStartDate
		// save session again
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
		}
		responseText = resp

	case "4":
		// event end date is the message sent
		eventEndDate := update.Message.Text
		_, _, _, err := utils.ParseDate(eventEndDate)
		if err != nil {
			responseText = utils.ErrorResponses["dateError"]
			userStateData.State = utils.UserStates["createEvent4"]
			SaveUserState(username, userStateData)
			break
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventEndDate = eventEndDate
		// save session again
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "5":
		// event start time is the message sent
		eventStartTime := update.Message.Text
		_, _, err := utils.ParseTime(eventStartTime)
		if err != nil {
			responseText = utils.ErrorResponses["timeError"]
			userStateData.State = utils.UserStates["createEvent5"]
			SaveUserState(username, userStateData)
			break
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventStartTime = eventStartTime
		// save session again
		resp, err := handleStateResponse(username, userStateData, event)
		responseText = resp

	case "6":
		// event end time is the message sent
		eventEndTime := update.Message.Text
		_, _, err := utils.ParseTime(eventEndTime)
		if err != nil {
			responseText = "You entered an invalid time, please try again and stick to the format HH:MM"
			userStateData.State = utils.UserStates["createEvent6"]
			SaveUserState(username, userStateData)
			break
		}

		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)

		event.EventEndTime = eventEndTime
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "7":
		// event timezone is the message sent
		eventTimeZone := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventTimeZone = eventTimeZone
		// save session again
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "8":
		// event location is the message sent
		eventLocation := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventLocation = eventLocation
		// save session again
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp

	case "9":
		// event attendees are the message sent or /retry
		var eventAttendees string
		if update.Message.Text == "/retry" {
			// get user's session ID and update the event associated to it
			sessionID = GetUserState(username).SessionID
			event := GetSession(sessionID)
			eventAttendees = strings.Join(event.EventAttendees, ",")
		} else {
			eventAttendees = update.Message.Text
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventAttendees = strings.Split(eventAttendees, ",")
		// Get event from Redis and create event on Google Calendar
		err := CreateCalendarEvent(userStateData.UserEmail, *event)
		// save session again
		if err != nil {
			log.Printf("An error occured while creating event: %v", err)
			responseText = utils.ErrorResponses["eventError"]
			break
		}
		resp, err := handleStateResponse(username, userStateData, event)
		if err != nil {
			log.Printf("An error occured: %v", err)
			break
		}
		responseText = resp
	}

	textMsg = tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
	textMsg.ReplyToMessageID = update.Message.MessageID

	return textMsg
}

func handleStateResponse(username string, stateData utils.UserStateData, event *utils.EventData) (string, error) {
	err := SaveSession(stateData.SessionID, event)
	if err != nil {
		return "", err
	}

	stateInt, _ := strconv.Atoi(stateData.State)
	var stateStr string
	if stateInt == 9 {
		stateStr = "eventCreated"
	} else {
		stateStr = fmt.Sprintf("createEvent%d", stateInt+1)
	}

	stateData.State = utils.UserStates[stateStr]
	SaveUserState(username, stateData)
	responseText := utils.StatefulResponseTemplates[stateStr]
	return responseText, nil
}

// func HandleHelp(update tgbotapi.Update) tgbotapi.MessageConfig {
// }
