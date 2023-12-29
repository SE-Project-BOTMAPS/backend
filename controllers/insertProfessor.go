package controllers

import (
	"encoding/json"
	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
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
		data := PrepareProfessorData()
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
func PrepareProfessorData() Professors {
	jsonFile, err := os.Open("professor.json")
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	content, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var proData Professors
	err = json.Unmarshal(content, &proData)
	if err != nil {
		panic(err)
	}

	return proData
}
