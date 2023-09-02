package utils

var ResponseTemplates = map[string]string{
	"start": `Hello👋, {{.username}}
I am your friendly 🗓️G-CalendarPal bot🤖
Kindly click the Signin button below to authorize and grant me access to your Google calendar`,
	"help": `Oh oh oh, Ebenezer your help has come...`,
}

var StatefulResponseTemplates = map[string]string{
	"createEvent1": `Okay, let's create a Google Calendar event for you...
What is the title of your event?`,
	"createEvent2": `Okay, what is the description of your event?`,
	"createEvent3": `Okay, what is the start date of your event?`,
	"createEvent4": `Okay, what is the end date of your event?`,
	"createEvent5": `Okay, what is the start time of your event?`,
	"createEvent6": `Okay, what is the end time of your event?`,
	"createEvent7": `Okay, what is the timezone of your event?`,
	"createEvent8": `Okay, what is the location of your event? Type "virtual" if it's a virtual event, and I will create a Google meet link for you.`,
	"createEvent9": `Now enter the emails of the attendees of your event. They will be notified via email.`,
	"eventCreated": `Congratulations🫣🎉🚀
Your event has been created successfully! 
An email has been sent out to you and your attendees also. Check your Google Calendar to view the event details.`,
}
