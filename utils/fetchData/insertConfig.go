package fetchData

import (
	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

func InsertConfig(configs Configuration, db *gorm.DB) {
	tx := db.Begin()
	tx.Exec("DELETE FROM configs")
	tx.Exec("ALTER TABLE configs AUTO_INCREMENT = 1")

	for _, subCalendar := range configs.SubCalendars.SubCalendars {
		config := models.Config{SubID: subCalendar.Id, Name: subCalendar.Name, Active: subCalendar.Active, Color: subCalendar.Color}
		tx.Create(&config)
	}
	tx.Commit()
}
