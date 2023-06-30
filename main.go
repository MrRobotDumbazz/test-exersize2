package main

import (
	"encoding/csv"
	"log"
	"os"
	"parser/models"

	"github.com/gocolly/colly"
)

func main() {
	thead := models.TableHead{}
	tbody := models.TableBody{}
	var tbodies []models.TableBody
	c := colly.NewCollector()
	c.OnHTML(".table", func(e *colly.HTMLElement) {
		row_rank := e.ChildText(".row__rank")
		row_about := e.ChildText(".row__about")
		row_category := e.ChildText(".row__category")
		row_subscribers := e.ChildText(".row__subscribers")
		row_audience := e.ChildText(".row__audience")
		row_authentic := e.ChildText(".row__authentic")
		Row_engagement := e.ChildText(".row__engagement")
		thead = models.TableHead{
			Row_rank:        row_rank,
			Row_about:       row_about,
			Row_category:    row_category,
			Row_subscribers: row_subscribers,
			Row_audience:    row_audience,
			Row_authentic:   row_authentic,
			Row_engagement:  Row_engagement,
		}
		log.Println(thead)
		rank := e.ChildText(".rank")
		contributor := e.ChildText(".contributor")
		category := e.ChildText(".category")
		subscribers := e.ChildText(".subscribers")
		audience := e.ChildText(".audience")
		authentic := e.ChildText(".authentic")
		engagement := e.ChildText(".engagement")
		tbody = models.TableBody{
			Rank:         rank,
			Contributtor: contributor,
			Category:     category,
			Subscribers:  subscribers,
			Audience:     audience,
			Authentic:    authentic,
			Engagement:   engagement,
		}
		tbodies = append(tbodies, tbody)
		log.Println(tbody)
	})
	err := c.Visit("https://hypeauditor.com/top-instagram-all-russia/")
	if err != nil {
		log.Println(err)
		return
	}
	file, err := os.Create("influencers.csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	headers := []string{
		"Rank",
		"About",
		"Category",
		"Subscribers",
		"Audience",
		"Authentic",
		"Engagement",
	}
	writer.Write(headers)
	defer writer.Flush()
	for _, tbody := range tbodies {
		record := []string{
			tbody.Rank,
			tbody.Contributtor,
			tbody.Category,
			tbody.Subscribers,
			tbody.Audience,
			tbody.Authentic,
			tbody.Engagement,
		}
		writer.Write(record)
	}
}
