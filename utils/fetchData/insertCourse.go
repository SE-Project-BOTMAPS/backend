package fetchData

import (
	"log"
	"strings"
	"time"

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

	dummyLocation := models.Location{Location: "DUMMYLOCx8zh2EjK1nqI9o3mXsLbNpVw"}
	tx.Create(&dummyLocation)
	for i, Event := range data.Event {
		location := models.Location{Location: Event.Location}
		professor := models.Professor{DataWho: Event.Who}
		tx.FirstOrCreate(&location, models.Location{Location: Event.Location})
		tx.FirstOrCreate(&professor, models.Professor{DataWho: Event.Who, OfficeLocationID: dummyLocation.ID})

		courses[i] = models.Course{
			Title:       Event.Title,
			StartTime:   Event.StartDt,
			EndTime:     Event.EndDt,
			LocationID:  location.ID,
			ProfessorID: professor.ID,
			Day:	 	 parseDay(Event),
		}
	}
	tx.Create(&courses)
	tx.Commit()
}

func parseDay(event Event) string {
	rrule := strings.Split(event.Rrule, ";")
	for _, instance := range rrule {
		mapping := strings.Split(instance, "=")
		if(mapping[0] == "BYDAY") {
			return getFullName(mapping[1])
		}
	}

	// In case there are no BEGIN instance
	startTime, err := time.Parse(time.RFC3339, event.StartDt)
	if(err != nil) {
		log.Print("No day provided for: " + event.Id)
		return ""
	}

	datePortion := startTime.Format("Monday")
	return datePortion

}

var dayReplacer = strings.NewReplacer(
	"MO", "Monday",
	"TU", "Tuesday",
	"WE", "Wednesday",
	"TH", "Thursday",
	"FR", "Friday",
	"SA", "Saturday",
	"SU", "Sunday",
)

func getFullName(days string) string {
	days = strings.ToUpper(days) 
	days = dayReplacer.Replace(days)

	return days
}