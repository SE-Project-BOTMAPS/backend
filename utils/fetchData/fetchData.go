package fetchData

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
{
            "id": "1609900170-rid-1704160800",
            "series_id": 1609900170,
            "remote_id": null,
            "subcalendar_id": 6820246,
            "subcalendar_ids": [
                6820246
            ],
            "all_day": false,
            "rrule": "FREQ=WEEKLY;UNTIL=20240308T235959+07:00",
            "title": "Project Meetings",
            "who": "dome",
            "location": "403",
            "notes": "",
            "version": "cb2239b1e737",
            "readonly": false,
            "tz": "Asia/Bangkok",
            "attachments": [],
            "start_dt": "2024-01-02T09:00:00+07:00",
            "end_dt": "2024-01-02T10:00:00+07:00",
            "ristart_dt": "2024-01-02T02:00:00+00:00",
            "rsstart_dt": "2023-12-05T09:00:00+07:00",
            "creation_dt": "2023-11-29T12:40:25+07:00",
            "update_dt": null,
            "delete_dt": null
        },*/

type Events struct {
	Event []Event `json:"events"`
}

type Event struct {
	Id       string `json:"id"`
	SubID    int    `json:"subcalendar_id"`
	Rrule    string `json:"rrule"`
	Title    string `json:"title"`
	Who      string `json:"who"`
	Location string `json:"location"`
	Notes    string `json:"notes"`
	StartDt  string `json:"start_dt"`
	EndDt    string `json:"end_dt"`
}

func FetchData(sDate, eDate string) (Events, error) {
	url := fmt.Sprintf("https://api.teamup.com/ksg7y4nwkfp7q6xyio/events?startDate=%s&endDate=%s", sDate, eDate)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request. ", err)
		return Events{}, err
	}

	req.Header.Set("Teamup-Token", os.Getenv("TEAMUP_KEY"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request. ", err)
		return Events{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error closing response body. ", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code: ", resp.StatusCode)
		return Events{}, err
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		return Events{}, err
	}

	var events Events
	err = json.Unmarshal(responseData, &events)
	if err != nil {
		log.Fatal("Error unmarshalling. ", err)
		return Events{}, err
	}
	return events, nil
}
