package utils

import (
	"testing"
)

func TestParseTimeSuite(t *testing.T) {
	t.Run("Test Valid Time", func(t *testing.T) {
		hour, min, err := ParseTime("12:15")
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if hour != 12 || min != 15 {
			t.Errorf("Expected 12:15, got %v:%v", hour, min)
		}
	})

	t.Run("Test Invalid Format", func(t *testing.T) {
		_, _, err := ParseTime("123456")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "invalid time format" {
			t.Errorf("Expected error message: Invalid time format, got: %v", err.Error())
		}

		_, _, err = ParseTime("12-15")
		if err == nil || err.Error() != "invalid time format" {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "invalid time format" {
			t.Errorf("Expected error message: Invalid time format, got: %v", err.Error())
		}
	})

	t.Run("Test Invalid Hour", func(t *testing.T) {
		_, _, err := ParseTime("25:30")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "hour is invalid, must be between 0 and 23" {
			t.Errorf("Expected error message: hour is invalid, must be between 0 and 23, got: %v", err.Error())
		}
	})

	t.Run("Test Invalid Minute", func(t *testing.T) {
		_, _, err := ParseTime("12:60")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "minute is invalid, must be between 0 and 59" {
			t.Errorf("Expected error message: minute is invalid, must be between 0 and 59, got: %v", err.Error())
		}
	})
}

func TestParseDateSuite(t *testing.T) {
	t.Run("Test Valid Date", func(t *testing.T) {
		year, month, day, err := ParseDate("2023-09-30")
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if year != 2023 || month != 9 || day != 30 {
			t.Errorf("Expected 2020-09-30, got %v-%v-%v", year, month, day)
		}
	})

	t.Run("Test invalid date format", func(t *testing.T) {
		y, m, d, err := ParseDate("2020-09-31-12")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err != nil && y != 0 && m != 0 && d != 0 {
			t.Errorf("Expected 0-0-0, got %v-%v-%v", y, m, d)
		}
	})

	t.Run("Test invalid date", func(t *testing.T) {
		_, _, _, err := ParseDate("2020-15-32")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
	})

	t.Run("Test invalid year", func(t *testing.T) {
		_, _, _, err := ParseDate("2019-09-30")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "year is invalid, must be greater than or equal to current year" {
			t.Errorf("Expected error message: year is invalid, must be greater than or equal to current year, got: %v", err.Error())
		}
	})

	t.Run("Test invalid month", func(t *testing.T) {
		_, _, _, err := ParseDate("2023-08-31")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "month is invalid, must be greater than or equal to current month" {
			t.Errorf("Expected error message: month is invalid, must be greater than or equal to current month, got: %v", err.Error())
		}
	})

	t.Run("Test invalid day", func(t *testing.T) {
		_, _, _, err := ParseDate("2023-09-03")
		if err == nil {
			t.Errorf("Expected error, but got none")
		}
		if err.Error() != "day is invalid, must be greater than or equal to current day" {
			t.Errorf("Expected error message: day is invalid, must be greater than or equal to current day, got: %v", err.Error())
		}
	})
}

func TestGenerateDateTime(t *testing.T) {
	t.Run("Test Valid Date and Time", func(t *testing.T) {
		dateTime, err := GenerateDateTime("2023-09-30", "12:15")
		if err != nil {
			t.Errorf("Expected none, got: %v", err)
		}
		if dateTime != "2023-9-30T12:15:00" {
			t.Errorf("Expected 2023-9-30T12:15:00, got %v", dateTime)
		}
	})

	t.Run("Test Invalid Date", func(t *testing.T) {
		dateTime, _ := GenerateDateTime("2023-09-31", "12:15")
		if dateTime != "0001-1-1T0:0:0" {
			t.Errorf("Expected 0001-1-1T0:0:0, got %v", dateTime)
		}
	})
}
