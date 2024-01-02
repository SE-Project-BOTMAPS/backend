package fetchData

import (
	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func InsertCourse(data Events, db *gorm.DB) {
	for _, Event := range data.Event {
		//	insert location to get foreign key to use in course
		location := models.Location{DataLocation: Event.Location}
		db.Create(&location)
		//	insert professor to get foreign key to use in course
		professor := models.Professor{DataWho: Event.Who}
		db.Create(&professor)
		//	insert course
		course := models.Course{Title: Event.Title, StartTime: Event.StartDt, EndTime: Event.EndDt, LocationID: location.ID, ProfessorID: professor.ID}
		db.Create(&course)
	}
}
