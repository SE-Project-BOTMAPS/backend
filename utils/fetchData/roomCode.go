package fetchData

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

type DayCourseMap map[string][]models.Course

func RoomCode(room_code string, db *gorm.DB) (DayCourseMap, error) {
	
	var locations []models.Location
	var courses []models.Course
	emptymap := DayCourseMap{}

	regexp := "%" + room_code + "%"

	// Query all locations with the room code
	err1 := db.Where("Location LIKE ?", regexp).Find(&locations).Error
	if len(locations) == 0 {
		message := "No such room found: " + room_code
		return emptymap, fmt.Errorf("%w: %s", gorm.ErrRecordNotFound, message)
	}
	if err1 != nil {
		return emptymap, err1
	}

	// Retrieve location IDs
	var locationIds []int64
	for _,location := range locations {
		locationIds = append(locationIds,location.ID)
	}
	
	// Query all courses with the location ID
	err2 := db.Preload("Location").Preload("Professor").Where("location_id IN ?", locationIds).Find(&courses).Error
	if err2 != nil {
		log.Println(err2)
		return emptymap, err2
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
		startTime, err := time.Parse("2006-01-02T15:04:05-07:00", course.StartTime)
		if err != nil {
			dayCourseMap["na"] = append(dayCourseMap["na"], course)
			continue 
		}

		key := daysMapping[startTime.Weekday().String()]
		dayCourseMap[key] = append(dayCourseMap[key], course)
	}

	// sorting by start time
	sortCoursesByStartTime := func(courses []models.Course) {
		sort.Slice(courses, func(i, j int) bool {
            timeI, _ := time.Parse("2006-01-02T15:04:05-07:00", courses[i].StartTime)
            timeJ, _ := time.Parse("2006-01-02T15:04:05-07:00", courses[j].StartTime)
            return timeI.Format("15:04:05") < timeJ.Format("15:04:05")
		})
	}

	for key, day := range dayCourseMap {
		if key == "na" {
			continue
		}

		sortCoursesByStartTime(day)
	}

	return dayCourseMap, nil
}