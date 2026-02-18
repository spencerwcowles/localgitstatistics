package main

import (
	"testing"
	"time"
)

func TestGetBeginningOfDay(t *testing.T) {
	now := time.Now()
	startOfDay := getBeginningOfDay(now)

	if startOfDay.Hour() != 0 || startOfDay.Minute() != 0 || startOfDay.Second() != 0 || startOfDay.Nanosecond() != 0 {
		t.Errorf("Expected start of day to be 00:00:00, got %v", startOfDay)
	}

	if startOfDay.Year() != now.Year() || startOfDay.Month() != now.Month() || startOfDay.Day() != now.Day() {
		t.Errorf("Expected start of day to be same date as now, got %v", startOfDay)
	}
}

func TestCountDaysSinceDate(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	fiveDaysAgo := now.AddDate(0, 0, -5)

	if countDaysSinceDate(yesterday) != 1 {
		t.Errorf("Expected 1 day since yesterday, got %d", countDaysSinceDate(yesterday))
	}

	if countDaysSinceDate(fiveDaysAgo) != 5 {
		t.Errorf("Expected 5 days since 5 days ago, got %d", countDaysSinceDate(fiveDaysAgo))
	}
}

func TestCalcOffset(t *testing.T) {
	offset := calcOffset()
	weekday := time.Now().Weekday()

	var expected int
	switch weekday {
	case time.Sunday:
		expected = 7
	case time.Monday:
		expected = 6
	case time.Tuesday:
		expected = 5
	case time.Wednesday:
		expected = 4
	case time.Thursday:
		expected = 3
	case time.Friday:
		expected = 2
	case time.Saturday:
		expected = 1
	}

	if offset != expected {
		t.Errorf("Expected offset to be %d for weekday %s, got %d", expected, weekday, offset)
	}
}
