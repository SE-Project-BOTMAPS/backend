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

	var offices []models.Office

	tx := db.Begin()
	tx.Exec("DELETE FROM offices")
	tx.Exec("ALTER TABLE offices AUTO_INCREMENT = 1")

	c.OnHTML("div.panel-boxtitle2", func(e *colly.HTMLElement) {
		firstname := strings.ToLower(strings.Fields(e.DOM.Find("font:nth-child(1)").Text())[0])
		office := e.DOM.Find("font:nth-child(8)").Text()

		var location models.Location
		var professor models.Professor
		tx.FirstOrCreate(&professor, models.Professor{DataWho: firstname})
		tx.FirstOrCreate(&location, models.Location{Location: office})

		offices = append(offices, models.Office{LocationID: location.ID, ProfessorID: professor.ID})
	})

	// Start the scraping process
	err := c.Visit(url)
	if err != nil {
		return err
	}

	tx.Create(&offices)
	tx.Commit()

	return nil
}