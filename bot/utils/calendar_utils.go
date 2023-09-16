package utils

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func ParseTime(time string) (int, int, error) {
	if len(time) != 5 {
		return 0, 0, errors.New("invalid time format")
	}

	// Check if the input string has the correct format "HH:MM"
	if time[2] != ':' {
		return 0, 0, errors.New("invalid time format")
	}

	hour := time[:2]
	min := time[3:5]

	hourInt, err := strconv.Atoi(hour)
	if err != nil || hourInt < 0 || hourInt > 23 {
		return 0, 0, errors.New("hour is invalid, must be between 0 and 23")
	}

	minInt, _ := strconv.Atoi(min)
	if err != nil || minInt < 0 || minInt > 59 {
		return 0, 0, errors.New("minute is invalid, must be between 0 and 59")
	}

	return hourInt, minInt, nil
}

func ParseDate(date string) (int, int, int, error) {
	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return 0, 0, 0, err
	}

	year := parsedDate.Year()
	month := int(parsedDate.Month())
	day := parsedDate.Day()

	// validate date
	currentDate := time.Now()
	currentYear := currentDate.Year()
	currentMonth := int(currentDate.Month())
	currentDay := currentDate.Day()

	if year < currentYear {
		return 0, 0, 0, errors.New("year is invalid, must be greater than or equal to current year")
	}

	if year >= currentYear && month < currentMonth {
		return 0, 0, 0, errors.New("month is invalid, must be greater than or equal to current month")
	}

	if year >= currentYear && month >= currentMonth && day < currentDay {
		return 0, 0, 0, errors.New("day is invalid, must be greater than or equal to current day")
	}

	return year, month, day, nil
}

// function to generate RFC3339 date-time string from date and time as entered by user
func GenerateDateTime(date, time string) (string, error) {
	year, month, day, err := ParseDate(date)

	if err != nil {
		log.Println("Error parsing date: ", err)
		return "", err
	}

	hour, min, err := ParseTime(time)
	if err != nil {
		log.Println("Error parsing time: ", err)
		return "", err
	}

	return generateRFC3339(year, month, day, hour, min), err
}

func generateRFC3339(year, month, day, hour, min int) string {
	dateTime := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day) + "T" + strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":00"
	log.Println("Datetime: ", dateTime)
	return dateTime
}
