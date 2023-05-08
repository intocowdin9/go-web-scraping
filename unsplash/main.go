package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("unsplash.com", "www.unsplash.com"))

	scrapeUrl := "https://unsplash.com/s/photos/palestine"

	file, err := os.Create("urls.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c.OnHTML(".mef9R", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildAttr("a", "href"))
		links := []string{h.ChildAttr("a", "href")}

		writer.Write(links)
	})

	c.OnHTML(".AYOsT", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText(".CwMIr DQBsa p1cWU jpBZ0 AYOsT Olora I0aPD dEcXu"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visit ", r.URL)
	})

	c.Visit(scrapeUrl)
}
