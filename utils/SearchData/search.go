package searchData

import (
	"log"
	"strings"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

type SearchData struct {
	Title string `json:"title"`
	Locations string `json:"locations"`
	Professor string `json:"professor"`
	Day string `json:"day"`
}

func Search(keyword string, db *gorm.DB) ([]models.Course, error){

	var professorIds []int64
	var roomIds []int64
	var courses []models.Course
	
	regexp := "%" + strings.ToLower(keyword) + "%"
	err1 := db.Table("professors").Where("LOWER(data_who) LIKE ? OR LOWER(full_name) LIKE ?", regexp, regexp).Select("professors.ID").Find(&professorIds).Error
	if(err1 != nil) {
		return nil,err1
	}
	log.Println(professorIds)

	err2 := db.Table("locations").Where("LOWER(Location) LIKE ?", regexp).Select("locations.ID").Find(&roomIds).Error
	if(err2 != nil) {
		return nil,err2
	}

	queryParam := "title LIKE ?"
	var args []interface{}
	args = append(args, regexp)
	
	if len(professorIds) > 0 {
		queryParam += " OR professor_id IN (?)"
		args = append(args, professorIds)
	}

	if len(queryParam) > 0 {
		queryParam += " OR location_id IN (?)"
		args = append(args, roomIds)
	}
	
	err3 := db.Preload("Professor").Preload("Location").Where(queryParam, args...).Find(&courses).Error
	if(err3 != nil) {
		return nil,err3
	}

	return courses, nil
}