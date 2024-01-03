package fetchData

import (
	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func InsertCourse(data Events, db *gorm.DB) {
	tx := db.Begin()
	//Empty database
	tx.Exec("DELETE FROM courses")
	tx.Exec("DELETE FROM locations")
	tx.Exec("DELETE FROM professors")
	//reset auto increment
	tx.Exec("ALTER TABLE courses AUTO_INCREMENT = 1")
	tx.Exec("ALTER TABLE locations AUTO_INCREMENT = 1")
	tx.Exec("ALTER TABLE professors AUTO_INCREMENT = 1")

	courses := make([]models.Course, len(data.Event))

	for i, Event := range data.Event {
		location := models.Location{Location: Event.Location}
		professor := models.Professor{DataWho: Event.Who}
		tx.FirstOrCreate(&location, models.Location{Location: Event.Location})
		tx.FirstOrCreate(&professor, models.Professor{DataWho: Event.Who})

		courses[i] = models.Course{
			Title:       Event.Title,
			StartTime:   Event.StartDt,
			EndTime:     Event.EndDt,
			LocationID:  location.ID,
			ProfessorID: professor.ID,
		}
	}
	tx.Create(&courses)
	tx.Commit()
}
