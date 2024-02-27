package fetchData

import (
	"log"
	"strconv"
	"sync"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"gorm.io/gorm"
)

type CourseTitle struct {
	ID    int64  `json:"id" `
	Title string `json:"title"`
}

type FullTitle struct {
	CourseTitleEng string `json:"course_title_eng"`
	CourseTitleTha string `json:"course_title_tha"`
}

func InsertCourseTitle(db *gorm.DB) error {

	// Get all courses code from db
	courseTitleMap, err1 := insertCoursesID(db)
	if(err1 != nil) {
		log.Println("Program fail to query all courses ID: ", err1)
		return err1
	}

	var wg sync.WaitGroup
	for courseID, courseFullTitle := range courseTitleMap {
		// Increment the wait group counter
		wg.Add(1)
		go func(cid int, cft FullTitle) {
			// Decrement the counter when the go routine completes
			defer wg.Done()

			var err1 error
			cft.CourseTitleEng, cft.CourseTitleTha, err1 = FetchCourseTitle(cid)
			if cft.CourseTitleEng == "" || err1 != nil { 
				return 
			}
			
			err2 := dbAssignCourseTitle(cid, cft, db)
			if err2 != nil {
				log.Println("Error assigning new title: ", err2)
			}
		}(courseID, courseFullTitle)
	}
	// Wait for all the checkWebsite calls to finish
	wg.Wait()

	return nil
}

func insertCoursesID(db *gorm.DB) (map[int]FullTitle, error) {
	var filteredCourses []CourseTitle
	titleMap := make(map[int]FullTitle)
	regexp := `^\d{6}.*` // 6 digits followed by any string
	
	err1 := db.Table("courses").Select("ID, Title").Where("REGEXP_LIKE(Title, ?)", regexp).Find(&filteredCourses).Error
	if(err1 != nil) {
		log.Println("Error querying courses from database: ", err1)
		return titleMap, err1
	}

	for _,filteredCourse := range filteredCourses {
		filteredCourse.Title = filteredCourse.Title[:6]
		courseID,err1 := strconv.Atoi(filteredCourse.Title)
		if(err1 != nil) {
			log.Println("Error at string to integer conversion: ", err1)
			return titleMap, err1
		}

		err2 := db.Table("courses").
				   Where("ID = ?", filteredCourse.ID).
				   Update("course_id",courseID).
				   Error
		if(err2 != nil) {
			log.Println("Error updating courseID: ", err2)
			return titleMap, err2
		}

		_, hasKey := titleMap[courseID]
		if(!hasKey) {
			titleMap[courseID] = FullTitle{"",""}
		}
	}

	return titleMap, nil
}

func dbAssignCourseTitle(courseID int, fullTitle FullTitle, db *gorm.DB) error {
	titleEng := fullTitle.CourseTitleEng
	titleTha := fullTitle.CourseTitleTha
	var courseTitle models.CourseTitle
	err := db.Table("course_titles").
			  Where("course_id = ?", courseID).
			  Assign(models.CourseTitle{FullTitleEng: titleEng, FullTitleTha: titleTha}).
			  FirstOrCreate(&courseTitle, models.CourseTitle{CourseID: courseID}).
			  Error

	return err
}