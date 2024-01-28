package fetchData

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Events struct {
	Event []Event `json:"events"`
}

type Event struct {
	Id       string `json:"id"`
	Rrule    string `json:"rrule"`
	Title    string `json:"title"`
	Who      string `json:"who"`
	Location string `json:"location"`
	Notes    string `json:"notes"`
	StartDt  string `json:"start_dt"`
	EndDt    string `json:"end_dt"`
}

type Configuration struct {
	SubCalendars SubCalendars `json:"configuration"`
}

type SubCalendars struct {
	SubCalendars []SubCalendar `json:"subcalendars"`
}

type SubCalendar struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Color  int64  `json:"color"`
}

func FetchImprove(url string, model interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request. ", err)
	}

	req.Header.Set("Teamup-Token", os.Getenv("TEAMUP_KEY"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request. ", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error closing response body. ", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code: ", resp.StatusCode)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	err = json.Unmarshal(responseData, &model)
	if err != nil {
		log.Fatal("Error unmarshalling. ", err)
	}
}

func FetchData(url string, token string) (Events, error) {
	var events Events
	return events, nil
}
