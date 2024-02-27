package controllers

import (
	"os"

	"github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"github.com/gin-gonic/gin"
)

func (db *DbController) UpdateData(c *gin.Context) {
	var events fetchData.Events
	baseUrl := os.Getenv("BASE_URL") + "events?startDate=2024-01-22&endDate=2024-01-28"
	fetchData.FetchImprove(baseUrl, &events)
	fetchData.InsertCourse(events, db.Database)
	go fetchData.InsertOffice(db.Database)
	go fetchData.InsertCourseTitle(db.Database)
	c.JSON(200, gin.H{
		"message": "Data updated",
	})
}
