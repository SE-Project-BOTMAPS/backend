package models

type Professor struct {
	ID       int64  `json:"id" orm:"auto" gorm:"primary_key;index"`
	DataWho  string `json:"data_who" orm:"size(128)" gorm:"unique"`
	FullName string `json:"full_name" orm:"size(128)"`
	course   []Course
}
