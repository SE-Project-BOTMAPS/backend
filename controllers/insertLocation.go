package controllers

import (
	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/SE-Project-BOTMAPS/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Locations struct {
	Locations []Location `json:"locations"`
}

type Location struct {
	DataLocation string `json:"data_location"`
	RoomCode     string `json:"room_code"`
	Detail       string `json:"detail"`
}

func (db *DbController) InsertLocation(c *gin.Context) {
	err := db.Database.Transaction(func(tx *gorm.DB) error {
		tx.Exec("DELETE FROM locations")
		var data Locations
		utils.ReadJsonFile("locations.json", &data)
		for i := range data.Locations {
			location := models.Location{DataLocation: data.Locations[i].DataLocation, RoomCode: data.Locations[i].RoomCode, Detail: data.Locations[i].Detail}
			err := tx.Create(&location).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
