package internals

import (
	"g_calendar_pal/bot/services"
	"g_calendar_pal/bot/utils"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleBot(bot *tgbotapi.BotAPI) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	var msg tgbotapi.MessageConfig
	var state utils.UserStateData

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		// log.Println(update.Message.From.UserName, update.Message.Text, update.Message.IsCommand())
		// if update type is not a message
		if update.Message == nil {
			continue
		}
		if update.Message.Text != "" {
			state = services.GetUserState(update.Message.From.UserName)
			if update.Message.IsCommand() {
				switch update.Message.Text {
				case "/start":
					msg = services.HandleStart(update)
				// case "/help":
				// 	msg = services.HandleHelp(update)
				case "/newevent":
					state.State = utils.UserStates["createEvent"]
					msg = services.HandleCreateEvent(update, state)

				case "/retry":
					state.State = utils.UserStates["createEvent9"]
					msg = services.HandleCreateEvent(update, state)
				}
			} else {
				stateInt, _ := strconv.Atoi(state.State)
				if stateInt >= 1 && stateInt <= 9 {
					msg = services.HandleCreateEvent(update, state)
				}
			}
		}
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("An error occured: %v", err)
		}
	}
}

// So I just thought of something I didn't think of before, I'm handling the multiple states of creating an event inside a handler called HandleCreate, but this handler is only called to handle the /createEvent command. The createEvent command is only handled when the incoming message is a command, now I don't know how to go about this, since all the remaining create event steps wont come with the /createEvent command in the text. What do you think?
