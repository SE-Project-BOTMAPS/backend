package fetchData

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func convertToDailyEvent(event Event, roomCode string) Event {
	StartTime, err := time.Parse(time.RFC3339, event.StartDt)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return Event{}
	}

	EndTime, err := time.Parse(time.RFC3339, event.EndDt)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return Event{}
	}

	DatePortion := StartTime.Format("Monday")
	timePortionStart := StartTime.Format("15:04")
	timePortionEnd := EndTime.Format("15:04")

	return Event{
		Id:       event.Id,
		Rrule:    DatePortion,
		Title:    event.Title,
		Who:      event.Who,
		Location: roomCode,
		Notes:    event.Notes,
		StartDt:  timePortionStart,
		EndDt:    timePortionEnd,
	}
}

func splitLocation(eventLocation string) []string {
	if strings.Contains(eventLocation, "/") {
		return strings.Split(eventLocation, "/")
	} else if strings.Contains(eventLocation, "-") {
		return strings.Split(eventLocation, "-")
	}
	return []string{eventLocation}
}

func DailyData(floor int) ([]Event, error) {
	events, err := FetchData("", "")
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	var result []Event

	for _, event := range events.Event {
		if locationMatchesFloor(event.Location, floor) {
			roomCodes := splitLocation(event.Location)

			for _, roomCode := range roomCodes {
				result = append(result, convertToDailyEvent(event, roomCode))
			}
		}
	}

	return result, nil
}

func locationMatchesFloor(location string, floor int) bool {
	return strings.HasPrefix(location, strconv.Itoa(floor))
}
