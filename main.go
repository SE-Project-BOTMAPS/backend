package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/SE-Project-BOTMAPS/backend/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	loadEnv()
	db := connectDB()
	Migration(db)

	router := gin.New()
	router.Use(cors.Default())
	api := router.Group("/api")
	routers.MainRouter(api, db)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func loadEnv() {
	if os.Getenv("ENV") == "prod" {
		log.Println("Running in production mode, using container environment variables")
	} else {
		log.Println("Loading .env file for local development")
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
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

func Migration(db *gorm.DB) {
	errCourse := db.AutoMigrate(&models.Course{})
	errLocation := db.AutoMigrate(&models.Location{})
	errProfessor := db.AutoMigrate(&models.Professor{})
	errConfig := db.AutoMigrate(&models.Config{})
	errOffice := db.AutoMigrate(&models.Office{})
	if errCourse != nil || errLocation != nil || errProfessor != nil || errConfig != nil || errOffice != nil {
		log.Fatalf("Error migrating database: %s", errCourse.Error())
	}
}
