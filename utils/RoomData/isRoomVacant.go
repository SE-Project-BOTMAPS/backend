package fetchData

import (
	"os"
	"slices"
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
	currentTime := time.Now()

	for _, reservation := range reservations {

		location := reservation.Location

		strCompoundLocation := strings.ReplaceAll(location,"-","/")
		compoundLocation := strings.Split(strCompoundLocation,"/")

		if(room_code != location && !slices.Contains(compoundLocation,room_code)) {
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
		if start.Before(currentTime) && end.After(currentTime) {
			return false, nil
		}
	}

	return true, nil
}