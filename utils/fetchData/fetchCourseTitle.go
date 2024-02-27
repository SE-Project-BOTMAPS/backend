package fetchData

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"
)

type CourseInformation struct {
	CourseID       string `json:"CourseID"`
	CourseTitleEng string `json:"CourseTitleEng"`
	CourseTitleTha  string `json:"CourseTitleTha"`
}

func FetchCourseTitle(courseId int) (string, string, error) {
	currentYearBE := time.Now().Year() + 543
	years := []int{currentYearBE, currentYearBE-1, currentYearBE-2, currentYearBE-3}
	terms := []int{1,2,3}

	courseTitleTha := ""
	courseTitleEng := ""
	var err error
	for _,year := range years {
		for _,term := range terms {
			courseTitleTha, courseTitleEng, err = fetchCourseTitle(courseId, year, term)
			if(err != nil) { return courseTitleEng, courseTitleTha, err }
			if(courseTitleTha != "") { return courseTitleTha, courseTitleEng, nil }
		}
	}
	
	// In case of not found. 
	return "", "", nil
}

func fetchCourseTitle(courseId int, year int, term int) (string, string, error) {
    baseURL := "https://mis-api.cmu.ac.th"
    resource := "/tqf/v1/course-template"
    params := url.Values{}
    params.Add("courseid", strconv.Itoa(courseId))
    params.Add("academicyear", strconv.Itoa(year))
    params.Add("academicterm", strconv.Itoa(term))

    u, err := url.ParseRequestURI(baseURL)
	if(err != nil) {
		log.Println("Error parsing URI. ", err)
		return "","",err
	}

    u.Path = resource
    u.RawQuery = params.Encode()
    urlStr := fmt.Sprintf("%v", u)

	var coursesInfo []CourseInformation
	FetchImprove(urlStr, &coursesInfo)

	if(len(coursesInfo) == 0) {
		return "","", nil
	}

	return coursesInfo[0].CourseTitleEng, coursesInfo[0].CourseTitleTha, nil
}