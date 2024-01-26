package fetchData

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
)

func InsertConfig(data Events, db *gorm.DB) {
	url := os.Getenv("BASE_URL") + "configuration"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request. ", err)
		return
	}

	req.Header.Set("Teamup-Token", os.Getenv("TEAMUP_KEY"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request. ", err)
		return
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
		return
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		return
	}

	var events Events
	err = json.Unmarshal(responseData, &events)
	if err != nil {
		log.Fatal("Error unmarshalling. ", err)
		return
	}

	tx := db.Begin()
}
