package model

type Professor struct {
	DataWho  string `json:"data_who" orm:"size(128)" gorm:"primary_key"`
	FullName string `json:"full_name" orm:"size(128)"`
	course   []Course
}
