package main

import (
	"flag"
	"fmt"

	"github.com/gocolly/colly"
)

type start struct {
	Name      string
	Photo     string
	JobTitle  string
	BirthDate string
	Bio       string
	TopMovies []movie
}

type movie struct {
	Title string
	Year  string
}

func main() {
	month := flag.Int("month", 1, "Month to fetch birthday for")
	day := flag.Int("day", 1, "Day to fetch bithdays for")
	flag.Parse()
	crawl(*month, *day)
}

func crawl(month int, day int) {

	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)
	infoCollector := c.Clone()
	c.OnHTML(".mode-detail", func(e *colly.HTMLElement) {
		profileUrl := e.ChildAttr("div.lister-item-image > a", "href")
		profileUrl = e.Request.AbsoluteURL(profileUrl)
		infoCollector.Visit(profileUrl)
	})

	startUrl := fmt.Sprintf("https://www.imdb.com/search/name/?birth_monthday=%d-%d", month, day)
	c.Visit(startUrl)
}
