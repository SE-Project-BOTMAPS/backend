package controllers

import (
	"fmt"
	"github.com/SE-Project-BOTMAPS/backend/utils/fetchData"
	"github.com/gin-gonic/gin"
	"os"
)

func (db *DbController) UpdateConfig(c *gin.Context) {
	var configs fetchData.Configuration
	baseUrl := os.Getenv("BASE_URL") + "configuration"
	fetchData.FetchImprove(baseUrl, &configs)
	fmt.Println(configs)
	fetchData.InsertConfig(configs, db.Database)
	c.JSON(200, gin.H{
		"message": "Config updated",
	})
}
