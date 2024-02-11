package controllers

import (
	"github.com/SE-Project-BOTMAPS/backend/utils/SearchData"
	"github.com/gin-gonic/gin"
)

func (db *DbController) SearchData(c *gin.Context){
	keyword := c.Param("keyword")
	search, err := SearchData.Search(keyword)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error Searching data.",
		})
	}
	c.JSON(200, search)
}