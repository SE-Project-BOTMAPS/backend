package fetchData

import (
	"fmt"
	"sort"
	"time"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

type DayCourseMap map[string][]models.Course

func RoomCode(room_code string, db *gorm.DB) (DayCourseMap, []models.Office, error) {
	
	var locations []models.Location
	var courses []models.Course
	emptymap := DayCourseMap{}
	emptyoffices := []models.Office{}

	regexp := "%" + room_code + "%"


	// Query all locations with the room code
	err1 := db.Where("Location LIKE ?", regexp).Find(&locations).Error
	if len(locations) == 0 {
		message := "No such room found: " + room_code
		return emptymap, emptyoffices, fmt.Errorf("%w: %s", gorm.ErrRecordNotFound, message)
	}
	if err1 != nil {
		return emptymap, emptyoffices, err1
	}

	// Retrieve location IDs
	var locationIds []int64
	for _,location := range locations {
		locationIds = append(locationIds,location.ID)
	}

	// Note: After this line, the searching location must exist.

	// Query the owner of the office
	var offices []models.Office
	err2 := db.Preload("Location").Where("location_id IN ?", locationIds).Find(&offices).Error
	if err2 != nil {
		offices = emptyoffices
	}
	
	// Query all courses with the location ID
	err3 := db.Preload("Location").Preload("Professor").Where("location_id IN ?", locationIds).Find(&courses).Error
	if err3 != nil {
		return emptymap, offices, nil
	}

	// categorizing
	dayCourseMap := DayCourseMap{
		"mon": []models.Course{},
		"tue": []models.Course{},
		"wed": []models.Course{},
		"thu": []models.Course{},
		"fri": []models.Course{},
		"sat": []models.Course{},
		"sun": []models.Course{},
		"na": []models.Course{},
	}

	daysMapping := map[string]string{
		"Monday":    "mon",
		"Tuesday":   "tue",
		"Wednesday": "wed",
		"Thursday":  "thu",
		"Friday":    "fri",
		"Saturday":  "sat",
		"Sunday":    "sun",
	}
	
	for _, course := range courses {
		startTime, err := time.Parse(time.RFC3339, course.StartTime)
		endTime, _ := time.Parse(time.RFC3339, course.EndTime)

		if err != nil {
			dayCourseMap["na"] = append(dayCourseMap["na"], course)
			continue 
		}

		key := daysMapping[startTime.Weekday().String()]

		course.StartTime = startTime.Format("15:04")
		course.EndTime = endTime.Format("15:04")

		dayCourseMap[key] = append(dayCourseMap[key], course)
	}

	// sorting by start time
	sortCoursesByStartTime := func(courses []models.Course) {
		sort.Slice(courses, func(i, j int) bool {
            return courses[i].StartTime < courses[j].StartTime
		})
	}

	for key, day := range dayCourseMap {
		if key == "na" {
			continue
		}

		sortCoursesByStartTime(day)
	}

	return dayCourseMap, offices, nil
}