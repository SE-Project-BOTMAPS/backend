package models

type Course struct {
	ID          int64     `json:"id" orm:"auto" gorm:"primary_key"`
	DataId      string    `json:"data_id" orm:"size(128)"`
	Title       string    `json:"title" orm:"size(128)"`
	Code        string    `json:"code" orm:"size(64)"`
	StartTime   string    `json:"start_time" orm:"size(64)"`
	EndTime     string    `json:"end_time" orm:"size(64)"`
	LocationID  int64     `json:"location_id"`
	Location    Location  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,references:DataLocation;"`
	ProfessorID int64     `json:"professor_id"`
	Professor   Professor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,references:DataWho;"`
}