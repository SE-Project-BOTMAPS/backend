package controllers

import "gorm.io/gorm"

type DbController struct {
	Database *gorm.DB
}
