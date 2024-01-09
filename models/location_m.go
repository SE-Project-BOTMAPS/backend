package models

type Location struct {
	ID       int64  `json:"id" orm:"auto" gorm:"primary_key;index"`
	Location string `json:"location" orm:"size(128)"`
	Detail   string `json:"detail" orm:"size(128)"`
	course   []Course
}
