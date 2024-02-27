package fetchData

import (
	"os"
	"slices"
	"strings"
	"time"

	fetchData "github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"gorm.io/gorm"
)

type Events = fetchData.Events
type Event = fetchData.Event

func IsRoomVacant(room_code string, db *gorm.DB) (bool, Event, error) {
	room_code = strings.TrimSpace(room_code)

	var events Events
	baseUrl := os.Getenv("BASE_URL") + "events"
	fetchData.FetchImprove(baseUrl, &events)

	isVacant, event, err := hasCurrentReservation(room_code, events.Event, db)
    
	if err != nil {
        return true, Event{}, err
    }

	return isVacant, event, nil
}

func hasCurrentReservation(room_code string, events []Event, db *gorm.DB) (bool,Event,error) {
	currentTime := time.Now()

	for _, event := range events {

		location := event.Location

		strCompoundLocation := strings.ReplaceAll(location,"-","/")
		compoundLocation := strings.Split(strCompoundLocation,"/")

		if(room_code != location && !slices.Contains(compoundLocation,room_code)) {
			continue
		}

		start, err := time.Parse(time.RFC3339, event.StartDt)
		if err != nil {
			return true, Event{}, err
		}
		end, err := time.Parse(time.RFC3339, event.EndDt)
		if err != nil {
			return true, Event{},err
		}
		if start.Before(currentTime) && end.After(currentTime) {
			return false, event, nil
		}
	}

	return true, Event{}, nil
}