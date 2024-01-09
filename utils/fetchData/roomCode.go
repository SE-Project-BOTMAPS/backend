package fetchData

import (
	"fmt"
	"log"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func RoomCode(room_code string, db *gorm.DB) ([]models.Course, error) {
	
	var locations []models.Location
	var courses []models.Course
	emptyCourse := []models.Course{}

	regexp := "%" + room_code + "%"

	// Query all locations with the room code
	err1 := db.Where("Location LIKE ?", regexp).Find(&locations).Error
	if len(locations) == 0 {
		message := "No such room found: " + room_code
		return emptyCourse, fmt.Errorf("%w: %s", gorm.ErrRecordNotFound, message)
	}
	if err1 != nil {
		return emptyCourse, err1
	}

	// Retrieve location IDs
	var locationIds []int64
	for _,location := range locations {
		locationIds = append(locationIds,location.ID)
	}
	
	// Query all courses with the location ID
	err2 := db.Preload("Location").Preload("Professor").Where("location_id IN ?", locationIds).Find(&courses).Order("start_time").Error
	if err2 != nil {
		log.Println(err2)
		return emptyCourse, err2
	}

	return courses, nil
}