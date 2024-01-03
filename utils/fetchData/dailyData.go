package fetchData

import (
	"log"
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

// type Events struct {
// 	Event []Event `json:"events"`
// }

// type Event struct {
// 	Id       string `json:"id"`
// 	Rrule    string `json:"rrule"`
// 	Title    string `json:"title"`
// 	Who      string `json:"who"`
// 	Location string `json:"location"`
// 	Notes    string `json:"notes"`
// 	StartDt  string `json:"start_dt"`
// 	EndDt    string `json:"end_dt"`
// }

func DailyData(floor int) (Events, error){
	events, err := FetchData("", "") 
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	for i, event := range events.Event {
		location := event.Location
		if int(location[0]) != floor{
			events.Event = append(events.Event[:i], events.Event[i+1:]...)
		}
	}
	return events, nil
}
