package models

type Config struct {
	ID     int64  `json:"id" orm:"auto" gorm:"primary_key"`
	SubID  int64  `json:"sub_id"`
	Name   string `json:"name" orm:"size(64)"`
	Active bool   `json:"active"`
	Color  int64  `json:"color"`
}
