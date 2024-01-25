package controllers

import (
	"github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (db *DbController) DailyData(c *gin.Context) {
	// Convert the parameter to an integer
	floor, err := strconv.Atoi(c.Param("floor"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid floor parameter. Please provide a valid integer.",
		})
		return
	}

	// Call the DailyData function with the parsed floor value
	data, err := fetchData.DailyData(floor, db.Database)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error fetching data.",
		})
		return
	}
	result := gin.H{"study": data[0], "reserve": data[1]}

	c.JSON(200, result)
	// Insert the data into the database
	// fetchData.InsertCourse(data, db.Database)
}
