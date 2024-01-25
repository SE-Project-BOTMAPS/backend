package fetchData

import (
	"fmt"
	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
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
		SubID:    event.SubID,
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
	if len(eventLocation) < 3 {
		return []string{eventLocation}
	}

	data := strings.FieldsFunc(eventLocation, func(r rune) bool {
		return r == '/' || r == '-'
	})

	for _, location := range data {
		if len(location) < 3 {
			return []string{eventLocation}
		}
	}

	return data
}

func DailyData(floor int, db *gorm.DB) ([][]Event, error) {
	events, err := FetchData("", "")
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	config := models.Config{Name: "Room reservations", Active: true}

	var resultDB models.Config
	if err := db.Where(&config).First(&resultDB).Error; err != nil {
		log.Fatal("Error querying database:", err)
	}

	subID := resultDB.SubID

	study := []Event{}
	reserve := []Event{}

	for _, event := range events.Event {
		if locationMatchesFloor(event.Location, floor) {
			roomCodes := splitLocation(event.Location)

			for _, roomCode := range roomCodes {
				if event.SubID == int(subID) {
					reserve = append(reserve, convertToDailyEvent(event, roomCode))
				} else {
					study = append(study, convertToDailyEvent(event, roomCode))
				}
			}
		}
	}

	return [][]Event{study, reserve}, nil
}

func locationMatchesFloor(location string, floor int) bool {
	return strings.HasPrefix(location, strconv.Itoa(floor))
}
