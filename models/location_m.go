package models

type Location struct {
	ID           int64  `json:"id" orm:"auto" gorm:"primary_key"`
	DataLocation string `json:"data_location" orm:"size(128)"`
	RoomCode     string `json:"room_code" orm:"size(128)"`
	Detail       string `json:"detail" orm:"size(128)"`
	course       []Course
}
