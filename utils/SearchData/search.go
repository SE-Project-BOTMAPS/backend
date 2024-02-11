package SearchData

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type SearchData struct {
	Title string `json:"title"`
	Locations string `json:"locations"`
	Professor string `json:"professor"`
	Day string `json:"day"`
}

func Search(keyword string) ([]SearchData, error){

	// log.Printf("Hello, %s", keyword)
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	//Create data source name 
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	rows, err := db.Query("(SELECT c.title ,l.location ,p.data_who, c.`day`  FROM courses c JOIN locations l ON c.location_id = l.id JOIN professors p ON p.id = c.id)")
	

	if err != nil {
        log.Fatal("Error querying database:", err)
    }
    defer rows.Close()

	SearchResults := []SearchData{}
	for rows.Next() {
		var s SearchData
		if err := rows.Scan(&s.Title, &s.Locations, &s.Professor, &s.Day); err != nil {
			log.Fatal("Error scanning rows:", err)
		}
		SearchResults = append(SearchResults, s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}
	log.Print(SearchResults)
	
	return SearchResults, nil
}