package controllers

import (
	"github.com/gin-gonic/gin"
)

type Courses struct {
	Courses []Course `json:"courses"`
}

type Course struct {
	DataId    string `json:"data_id"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	Professor string `json:"professor"`
	StartDt   string `json:"start_dt"`
	EndDt     string `json:"end_dt"`
}

func (db *DbController) InsertCourse(c *gin.Context) {
	//err := db.Database.Transaction(func(tx *gorm.DB) error {
	//	tx.Exec("DELETE FROM courses")
	//	var data Courses
	//	utils.ReadJsonFile("courses.json", &data)
	//	for i := range data.Courses {
	//		course := models.Course{DataId: data.Courses[i].DataId, Title: data.Courses[i].Title, LocationID: data.Courses[i].Location, ProfessorID: data.Courses[i].Professor, StartTime: data.Courses[i].StartDt, EndTime: data.Courses[i].EndDt}
	//		err := tx.Create(&course).Error
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return
	//}
	c.JSON(200, gin.H{"message": "Success"})
}
