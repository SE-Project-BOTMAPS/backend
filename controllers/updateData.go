package controllers

import (
	"github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"github.com/gin-gonic/gin"
)

func (db *DbController) UpdateData(c *gin.Context) {
	data, err := fetchData.FetchData("2024-01-22", "2024-01-28")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error fetching data.",
		})
		return
	}
	fetchData.InsertCourse(data, db.Database)
}
