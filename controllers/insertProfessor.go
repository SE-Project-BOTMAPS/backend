package controllers

import (
	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/SE-Project-BOTMAPS/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Professors struct {
	Professors []Professor `json:"professors"`
}

type Professor struct {
	Who      string `json:"who"`
	FullName string `json:"fullName"`
}

func (db *DbController) InsertProfessor(c *gin.Context) {
	err := db.Database.Transaction(func(tx *gorm.DB) error {
		tx.Exec("DELETE FROM professors")
		var data Professors
		utils.ReadJsonFile("professor.json", &data)
		for i := range data.Professors {
			professor := models.Professor{DataWho: data.Professors[i].Who, FullName: data.Professors[i].FullName}
			err := tx.Create(&professor).Error
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
