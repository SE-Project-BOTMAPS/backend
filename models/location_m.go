package models

type Location struct {
	DataLocation string `json:"data_location" orm:"size(128)" gorm:"primary_key"`
	RoomCode     string `json:"room_code" orm:"size(128)"`
	Detail       string `json:"detail" orm:"size(128)"`
	course       []Course
}
