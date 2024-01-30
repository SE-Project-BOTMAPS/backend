package models

type Office struct {
	ID          int64     `json:"id" orm:"auto" gorm:"primary_key"`
	LocationID  int64     `json:"location_id" gorm:"index"`
	Location    Location  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProfessorID int64     `json:"professor_id" gorm:"index"`
	Professor   Professor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
