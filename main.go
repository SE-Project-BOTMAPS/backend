package main

import (
	"fmt"
	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/SE-Project-BOTMAPS/backend/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	loadEnv()
	db := connectDB()
	migration(db)

	router := gin.New()
	api := router.Group("/api")
	routers.MainRouter(api, db)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectDB() *gorm.DB {
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	DbUser := os.Getenv("DB_USER")
	DbPort := os.Getenv("DB_PORT")
	DbPassword := os.Getenv("DB_PASSWORD")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}
	return db
}

func migration(db *gorm.DB) {
	errCourse := db.AutoMigrate(&models.Course{})
	errLocation := db.AutoMigrate(&models.Location{})
	errProfessor := db.AutoMigrate(&models.Professor{})
	if errCourse != nil || errLocation != nil || errProfessor != nil {
		log.Fatalf("Error migrating database: %s", errCourse.Error())
	}
}
