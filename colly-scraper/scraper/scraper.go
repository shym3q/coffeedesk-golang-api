package scraper

import (
	"github.com/colly-scraper/models"
	"github.com/gocolly/colly"
)

type Scraper struct {
	AllowedDomains string
	ScrapedData    []models.Coffee
}

func (scraper *Scraper) Scrape() {
	c := colly.NewCollector(colly.AllowedDomains(scraper.AllowedDomains))

	finished := false

	c.OnHTML(".page-next.disabled", func(e *colly.HTMLElement) {
		finished = true
	})

	c.OnHTML("div.cms-listing-col", func(e *colly.HTMLElement) {
		if !finished {
			name := e.ChildAttr("a", "title")
			description := e.ChildText(".product-description")
			link := e.ChildAttr("a", "href")
			image := e.ChildAttr("img", "src")
			coffee := models.Coffee{Name: name, Description: description, Link: link, Image: image}
			if !(image == "") {
				scraper.ScrapedData = append(scraper.ScrapedData, coffee)
			}
		}
	})

	c.OnHTML(".page-next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL("?order=dostepnosc&p=" + e.ChildAttr("input", "value"))
		err := c.Visit(nextPage)
		if err != nil {
			return
		}
	})

	err := c.Visit("https://www.coffeedesk.pl/kawa/metoda-parzenia/przelewowe-metody-parzenia/")
	if err != nil {
		return
	}
}
