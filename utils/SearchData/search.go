package searchData

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func Search(keyword string, db *gorm.DB) ([]models.Course, error){

	regexp := "%" + strings.ToLower(keyword) + "%"

	professorClause := "professor_id IN (SELECT ID FROM professors WHERE LOWER(data_who) LIKE ? OR LOWER(full_name) LIKE ? OR native_name LIKE ?)"
	locationClause := "location_id IN (SELECT ID FROM locations WHERE LOWER(Location) LIKE ?)"
	courseIDClause := "course_id IN (SELECT course_id FROM course_titles WHERE LOWER(full_title_eng) LIKE ? OR full_title_tha LIKE ?)"
	queryParameters := fmt.Sprintf("title LIKE ? OR %s OR %s OR %s ",professorClause,locationClause,courseIDClause)
	args := []interface{}{regexp, regexp, regexp, regexp, regexp, regexp, regexp}

	// Query for courses matching the keyword and related entities
	var courses []models.Course
	err := db.Order("start_time").
			Preload("Professor").
			Preload("Location").
			Where(queryParameters, args...).
			Find(&courses).Error
	if err != nil {
		log.Println("Error querying selected courses: ", err)
		return nil, err
	}

	for i,course := range(courses) {
		StartTime, err := time.Parse(time.RFC3339, course.StartTime)
		if err != nil {
			log.Println("Error parsing start time:", err)
		}

		EndTime, err := time.Parse(time.RFC3339, course.EndTime)
		if err != nil {
			log.Println("Error parsing end time:", err)
		}
		
		day := StartTime.Weekday().String()

		courses[i].StartTime = StartTime.Format("15:04")
		courses[i].EndTime = EndTime.Format("15:04")
		courses[i].Day = day
	}

	return courses, nil
}