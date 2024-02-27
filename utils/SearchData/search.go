package searchData

import (
	"fmt"
	"strings"
	"time"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func Search(keyword string, db *gorm.DB) ([]models.Course, error){

	var professorIds []int64
	var roomIds []int64
	var courseIds []int
	var courses []models.Course
	
	regexp := "%" + strings.ToLower(keyword) + "%"

	err1 := db.Table("professors").
			   Where("LOWER(data_who) LIKE ? OR LOWER(full_name) LIKE ? OR native_name LIKE ?", regexp, regexp, regexp).
			   Select("professors.ID").
			   Find(&professorIds).Error
	if(err1 != nil) {
		return nil,err1
	}

	err2 := db.Table("locations").Where("LOWER(Location) LIKE ?", regexp).Select("locations.ID").Find(&roomIds).Error
	if(err2 != nil) {
		return nil,err2
	}

	err3 := db.Table("course_titles").
			   Where("LOWER(full_title_eng) LIKE ? OR full_title_tha LIKE ?", regexp, regexp).
			   Select("course_titles.course_id").
			   Find(&courseIds).Error
	if(err3 != nil) {
		return nil,err3
	}

	queryParam := "title LIKE ?"
	var args []interface{}
	args = append(args, regexp)
	
	if len(professorIds) > 0 {
		queryParam += " OR professor_id IN (?)"
		args = append(args, professorIds)
	}

	if len(roomIds) > 0 {
		queryParam += " OR location_id IN (?)"
		args = append(args, roomIds)
	}
	
	if len(courseIds) > 0 {
		queryParam += " OR course_id IN (?)"
		args = append(args, courseIds)
	}

	err4 := db.Order("start_time").Preload("Professor").Preload("Location").Where(queryParam, args...).Find(&courses).Error
	if(err4 != nil) {
		return nil,err4
	}

	for i,course := range(courses) {
		StartTime, err := time.Parse(time.RFC3339, course.StartTime)
		if err != nil {
			fmt.Println("Error parsing start time:", err)
		}

		EndTime, err := time.Parse(time.RFC3339, course.EndTime)
		if err != nil {
			fmt.Println("Error parsing end time:", err)
		}
		
		day := StartTime.Weekday().String()

		courses[i].StartTime = StartTime.Format("15:04")
		courses[i].EndTime = EndTime.Format("15:04")
		courses[i].Day = day
	}

	return courses, nil
}