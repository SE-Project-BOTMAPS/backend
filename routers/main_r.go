package routers

import (
	"github.com/SE-Project-BOTMAPS/backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRouter(router *gin.RouterGroup, db *gorm.DB) {
	ctrl := controllers.DbController{Database: db}

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.POST("/Professor", ctrl.InsertProfessor)
	router.POST("/Location", ctrl.InsertLocation)
	router.POST("/Course", ctrl.InsertCourse)
	router.POST("/data", ctrl.UpdateData)
}
