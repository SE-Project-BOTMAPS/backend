package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}
