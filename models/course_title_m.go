package models

type CourseTitle struct {
	CourseID     int    `json:"course_id" orm:"auto" gorm:"primary_key;index"`
	FullTitleTha string `json:"full_title_tha" orm:"size(128)"`
	FullTitleEng string `json:"full_title_eng" orm:"size(128)"`
}
