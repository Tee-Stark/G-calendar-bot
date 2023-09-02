package utils

import (
	"log"
	"strconv"
)

func ParseTime(time string) (int, int) {
	hour := time[:2]
	min := time[3:5]

	hourInt, _ := strconv.Atoi(hour)
	minInt, _ := strconv.Atoi(min)

	return hourInt, minInt
}

func ParseDate(date string) (int, int, int) {
	// log.Println("Date: ", date)
	year := date[:4]
	month := date[5:7]
	day := date[8:10]

	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := strconv.Atoi(month)
	dayInt, _ := strconv.Atoi(day)

	return yearInt, monthInt, dayInt
}

// function to generate RFC3339 date-time string from date and time as entered by user
func GenerateDateTime(date, time string) string {
	year, month, day := ParseDate(date)
	hour, min := ParseTime(time)

	return generateRFC3339(year, month, day, hour, min)
}

func generateRFC3339(year, month, day, hour, min int) string {
	dateTime := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day) + "T" + strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":00"
	log.Println("Datetime: ", dateTime)
	return dateTime
}
