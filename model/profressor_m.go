package model

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	DataWho  string `json:"data_who" orm:"size(128)"`
	FullName string `json:"full_name" orm:"size(128)"`
	course   []Course
}
