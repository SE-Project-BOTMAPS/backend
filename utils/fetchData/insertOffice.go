package fetchData

import (
	"strings"

	"github.com/SE-Project-BOTMAPS/backend/models"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
)

func InsertOffice(db *gorm.DB) error {

	url := "https://www.cpe.eng.cmu.ac.th/lecturer-thai.php"
	c := colly.NewCollector()

	tx := db.Begin()

    c.OnHTML("div.panel-people", func(e *colly.HTMLElement) {

        ce1 := e.DOM.Find("div.panel-boxtitle1")
        ce2 := e.DOM.Find("div.panel-boxtitle2")
		
        nativename := ce1.Find("a").Text()
        fullname := ce2.Find("font:nth-child(1)").Text()
        firstname := strings.ToLower(strings.Fields(fullname)[0])
        office := ce2.Find("font:nth-child(8)").Text()

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

	tx.Commit()

	return nil
}