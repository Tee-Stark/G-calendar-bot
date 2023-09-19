package utils

var ResponseTemplates = map[string]string{
	"start": `Hello👋, {{.username}}
I am G-CalendarPal bot🤖,
your friendly Google Calendar Pal🗓️
Kindly click the Signin button below to authorize and grant me access to your Google calendar`,
	"help": `Hey! I'm the easiest bot out there to use.
	Use /start to start a conversation with me and signin with Google.
	Use /newevent to begin the process of creating a new event, then respond to my prompts with your event details to create your event.
	Use /help to see this message again.`,
}

var ErrorResponses = map[string]string{
	"dateError":  "You entered an invalid date, please try again and stick to the format YYYY-MM-DD",
	"timeError":  "You entered an invalid time, please try again and stick to the format HH:MM",
	"eventError": "An error occured while creating event. Please try again later with the /retry command.",
}

var StatefulResponseTemplates = map[string]string{
	"createEvent1": `Okay, let's create a Google Calendar event for you...
What is the title of your event?`,
	"createEvent2": `Okay, what is the description of your event?`,
	"createEvent3": `When is your event starting? 
	Enter date in this format => YYYY-MM-DD`,
	"createEvent4": `When does your event end? 
	Enter date in this format => YYYY-MM-DD`,
	"createEvent5": `Alright, now tell me the starting time for your event.
	Enter time in this format => HH:MM (24hr format)`,
	"createEvent6": `Okay, what time does it end?
	Enter time in this format => HH:MM (24hr format)`,
	"createEvent7":  `I'll be using GMT+1 as the timezone for your event, is that okay?`,
	"createEvent8":  `Is your event virtual?.`,
	"createEvent8b": `Great! I'll be sure to create a Google meet link for your event.`,
	"createEvent8c": `Oh, okay. Kindly enter the location of your event.`,
	"createEvent9": `One last thing, enter the email of your attendees.
	Seperate with a comma ",", e.g me@gmail.com,you@gmail.com).`,
	"eventCreated": `Congratulations🫣🎉🚀
Your event has been created successfully! 
An email has been sent out to you and your attendees. Check your Google Calendar to view the event details.`,
}
