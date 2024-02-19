package fetchData

import (
	"log"
	"strings"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
)

func InsertOffice(db *gorm.DB) error {

	url := "https://www.cpe.eng.cmu.ac.th/lecturer-thai.php"
	c := colly.NewCollector()


	nativeNameMap := make(map[string]string)
	FullNameMap := make(map[string]string)

	tx := db.Begin()

    c.OnHTML("div.panel-people", func(e *colly.HTMLElement) {

        ce1 := e.DOM.Find("div.panel-boxtitle1")
        ce2 := e.DOM.Find("div.panel-boxtitle2")
		
        nativename := ce1.Find("a").Text()
        fullname := ce2.Find("font:nth-child(1)").Text()
        firstname := strings.ToLower(strings.Fields(fullname)[0])
        office := ce2.Find("font:nth-child(8)").Text()

		nativeNameMap[firstname] = nativename
		FullNameMap[firstname] = fullname

        var location models.Location
        var professor models.Professor
        tx.FirstOrCreate(&location, models.Location{Location: office})
        tx.Where(models.Professor{DataWho: firstname}).Assign(models.Professor{FullName: fullname, NativeName: nativename, OfficeLocationID:location.ID}).FirstOrCreate(&professor, models.Professor{DataWho: firstname})
    })

	// Start the scraping process
	err := c.Visit(url)
	if err != nil {
		return err
	}

	// Handle compound professors
	var compoundProfessors []models.Professor
	err1 := db.Where("data_who LIKE ?", "%,%").Find(&compoundProfessors).Error
	if(err1 != nil) {
		log.Panicln("Cannot query compound professors: " + err1.Error()) 
	}

	for _, cmpf := range(compoundProfessors) {
		compoundFirstNames := strings.Split(cmpf.DataWho, ",")
		compoundNativeName := ""
		compoundFullName := ""
		for _,firstName := range(compoundFirstNames) {
			compoundNativeName += nativeNameMap[firstName] + ","
			compoundFullName += FullNameMap[firstName] + ","
		}

		tx.Where(models.Professor{DataWho: cmpf.DataWho}).Updates(models.Professor{FullName: compoundFullName, NativeName: compoundNativeName})
	}

	tx.Commit()

	return nil
}