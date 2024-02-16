package controllers

import (
	"errors"

	"github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (db *DbController) RoomCode(c *gin.Context) {
	
	room_code := c.Param("room_code")

	// Call the RoomCode function with the room_code parameter
	dayCourseMap, officeOf, err := fetchData.RoomCode(room_code, db.Database)
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": "No such room found: " + room_code,
		})
		return
	}
	
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error fetching data.",
		})
		return
	}

	c.JSON(200, gin.H{"events": dayCourseMap, "officeOf": officeOf})
}