package services

import (
	"g_calendar_pal/bot/utils"
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
	case utils.UserStates["createEvent"]:
		// create new session and update userStateData
		sessionID = utils.GenerateShortID()
		userStateData.State = utils.UserStates["createEvent1"]
		userStateData.SessionID = sessionID
		// save session and user state

		SaveUserState(username, userStateData)
		// err := SaveSession(sessionID, event)
		// if err != nil {
		// 	log.Printf("An error occured while saving session: %v", err)
		// }
		responseText = utils.StatefulResponseTemplates["createEvent1"]

	case utils.UserStates["createEvent1"]:
		// event name is the message sent
		eventName := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		log.Println("Session ID: ", sessionID)
		event := utils.NewEvent() // &utils.EventData{}
		log.Println("Event: ", event)
		event.EventName = eventName
		// save session and state again
		err := SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
			break
		}
		userStateData.State = utils.UserStates["createEvent2"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["createEvent2"]

	case utils.UserStates["createEvent2"]:
		// event description is the message sent
		eventDesc := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		log.Println("Session ID: ", sessionID)
		event := GetSession(sessionID)
		log.Println("Event: ", event)
		event.EventDescription = eventDesc
		// save session and state again
		userStateData.State = utils.UserStates["createEvent3"]
		SaveUserState(username, userStateData)
		err := SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		responseText = utils.StatefulResponseTemplates["createEvent3"]

	case utils.UserStates["createEvent3"]:
		// event start date is the message sent
		eventStartDate := update.Message.Text
		_, _, _, err := utils.ParseDate(eventStartDate)
		if err != nil {
			responseText = "You entered an invalid date, please try again and stick to the format YYYY-MM-DD"
			userStateData.State = utils.UserStates["createEvent3"]
			SaveUserState(username, userStateData)
			break
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventStartDate = eventStartDate
		// save session again
		err = SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		userStateData.State = utils.UserStates["createEvent4"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["createEvent4"]

	case utils.UserStates["createEvent4"]:
		// event end date is the message sent
		eventEndDate := update.Message.Text
		_, _, _, err := utils.ParseDate(eventEndDate)
		if err != nil {
			responseText = "You entered an invalid date, please try again and stick to the format YYYY-MM-DD"
			userStateData.State = utils.UserStates["createEvent4"]
			SaveUserState(username, userStateData)
			break
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventEndDate = eventEndDate
		// save session again
		err = SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		userStateData.State = utils.UserStates["createEvent5"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["createEvent5"]

	case utils.UserStates["createEvent5"]:
		// event start time is the message sent
		eventStartTime := update.Message.Text
		_, _, err := utils.ParseTime(eventStartTime)
		if err != nil {
			responseText = "You entered an invalid time, please try again and stick to the format HH:MM"
			userStateData.State = utils.UserStates["createEvent5"]
			SaveUserState(username, userStateData)
			break
		}
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventStartTime = eventStartTime
		// save session again
		err = SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		userStateData.State = utils.UserStates["createEvent6"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["createEvent6"]

	case utils.UserStates["createEvent6"]:
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
		err = SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}

		userStateData.State = utils.UserStates["createEvent7"]
		SaveUserState(username, userStateData)

		responseText = utils.StatefulResponseTemplates["createEvent7"]

	case utils.UserStates["createEvent7"]:
		// event timezone is the message sent
		eventTimeZone := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventTimeZone = eventTimeZone
		// save session again
		err := SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		userStateData.State = utils.UserStates["createEvent8"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["createEvent8"]

	case utils.UserStates["createEvent8"]:
		// event location is the message sent
		eventLocation := update.Message.Text
		// get user's session ID and update the event associated to it
		sessionID = GetUserState(username).SessionID
		event := GetSession(sessionID)
		event.EventLocation = eventLocation
		// save session again
		err := SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		userStateData.State = utils.UserStates["createEvent9"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["createEvent9"]

	case utils.UserStates["createEvent9"]:
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
		// save session again
		err := SaveSession(sessionID, event)
		if err != nil {
			log.Printf("An error occured while saving session: %v", err)
		}
		// Get event from Redis and create event on Google Calendar
		// event = GetSession(sessionID)
		log.Println("Event log", event)
		err = CreateCalendarEvent(userStateData.UserEmail, *event)
		if err != nil {
			log.Printf("An error occured while creating event: %v", err)
			responseText = "An error occured while creating event. Please try again later with the /retry command."
			break
		}
		userStateData.State = utils.UserStates["eventCreated"]
		SaveUserState(username, userStateData)
		responseText = utils.StatefulResponseTemplates["eventCreated"]
	}

	textMsg = tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
	textMsg.ReplyToMessageID = update.Message.MessageID

	return textMsg
}

// func HandleHelp(update tgbotapi.Update) tgbotapi.MessageConfig {
// }
