package controllers

import (
	"github.com/SE-Project-BOTMAPS/backend/utils/searchData"
	"github.com/gin-gonic/gin"
)

func (db *DbController) SearchData(c *gin.Context){
	keyword := c.Param("keyword")
	
	courses, err := searchData.Search(keyword, db.Database)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error Searching data.",
		})
		return
	}
	c.JSON(200, courses)
}