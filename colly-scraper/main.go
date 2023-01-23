package main

import (
	"github.com/colly-scraper/handlers"
	"github.com/colly-scraper/scraper"
)

func main() {
	lastData := handlers.GetAll()
	s := scraper.Scraper{AllowedDomains: "www.coffeedesk.pl"}
	s.Scrape()
	for _, lastCoffee := range lastData {
		upToDate := false
		for _, newCoffee := range s.ScrapedData {
			if lastCoffee.Name == newCoffee.Name {
				upToDate = true
				break
			}
		}
		if !upToDate {
			handlers.Delete(lastCoffee)
		}
	}
	for _, coffee := range s.ScrapedData {
		handlers.Add(coffee)
	}
}
