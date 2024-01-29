package models

type Office struct {
	ID         int64    `json:"id" orm:"auto" gorm:"primary_key;index"`
	Professor  string   `json:"professor" orm:"size(128)"`
	LocationID int64    `json:"location_id" gorm:"index"`
	Location   Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
