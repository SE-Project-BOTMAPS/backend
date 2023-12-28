package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	DataLocation string `json:"data_location" orm:"size(128)"`
	RoomCode     string `json:"room_code" orm:"size(128)"`
	Detail       string `json:"detail" orm:"size(128)"`
	course       []Course
}
