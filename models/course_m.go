package models

type Course struct {
	ID          int64     `json:"id" orm:"auto" gorm:"primary_key"`
	DataId      string    `json:"data_id" orm:"size(128)"`
	Title       string    `json:"title" orm:"size(128)"`
	Code        string    `json:"code" orm:"size(64)"`
	Day         string    `json:"day" orm:"size(64)"`
	StartTime   string    `json:"start_time" orm:"size(64)"`
	EndTime     string    `json:"end_time" orm:"size(64)"`
	LocationID  int64     `json:"location_id" gorm:"index"`
	Location    Location  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProfessorID int64     `json:"professor_id" gorm:"index"`
	Professor   Professor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
