package fetchData

import (
	"log"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func RoomCode(room_code string, db *gorm.DB) ([]models.Course, error) {
	
	var locations []models.Location
	var courses []models.Course

	regexp := "%" + room_code + "%"

	// Query all locations with the room code
	err1 := db.Where("Location LIKE ?", regexp).Find(&locations).Error
	if err1 != nil {
		log.Println(err1)
		return []models.Course{}, err1
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
		return []models.Course{}, err2
	}

	return courses, nil
}