package models

type Office struct {
	ID        int64  `json:"id" orm:"auto" gorm:"primary_key;index"`
	Professor string `json:"professor" orm:"size(128)"`
	Location  string `json:"location" orm:"size(128)"`
}
