package fetchData

import (
	"os"
	"strings"
	"time"

	fetchData "github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"gorm.io/gorm"
)

type Reservations struct {
	Reservation []Reservation `json:"events"`
}

type Reservation struct {
	Id       string `json:"id"`
	Location string `json:"location"`
	StartDt  string `json:"start_dt"`
	EndDt    string `json:"end_dt"`
}

func IsRoomVacant(room_code string, db *gorm.DB) (bool, error) {
	room_code = strings.TrimSpace(room_code)

	var reservations Reservations
	baseUrl := os.Getenv("BASE_URL") + "events"
	fetchData.FetchImprove(baseUrl, &reservations)

	isVacant, err := hasCurrentReservation(room_code, reservations.Reservation)
    
	if err != nil {
        return true, err
    }

	return isVacant, nil
}

func hasCurrentReservation(room_code string, reservations []Reservation) (bool,error) {
	currentTime , _ := time.Parse(time.RFC3339, "2024-02-24T08:30:00+07:00")

	for _, reservation := range reservations {

		if(room_code != reservation.Location) {
			continue
		}

		start, err := time.Parse(time.RFC3339, reservation.StartDt)
		if err != nil {
			return true, err
		}
		end, err := time.Parse(time.RFC3339, reservation.EndDt)
		if err != nil {
			return true, err
		}
		if start.Before(currentTime) && end.After(currentTime) && reservation.Location == room_code {
			return false, nil
		}
	}

	return true, nil
}