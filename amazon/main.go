package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector(
		// Visit only the domain in the start URL
		colly.AllowedDomains("www.amazon.in"),
	)

	// Slice to hold the data
	var data [][]string

	// On every product row
	c.OnHTML("div[data-component-type='s-search-result']", func(e *colly.HTMLElement) {
		// Extract the title, price, and image URL
		title := strings.TrimSpace(e.ChildText("h2 a"))
		price := strings.TrimSpace(e.ChildText("span.a-price-whole"))
		image := e.ChildAttr("img.s-image", "src")

		// Convert price to a float and format it
		// priceFloat, _ := strconv.ParseFloat(price, 64)
		// priceFormatted := fmt.Sprintf("%.2f", priceFloat)

		// Append the data to the slice
		data = append(data, []string{title, price, image})
	})

	// On pagination link click
	c.OnHTML("ul.a-pagination li.a-last a", func(e *colly.HTMLElement) {
		// Visit the next page
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Start the collector on the search results page
	c.Visit("https://www.amazon.in/s?k=server&crid=1V4XZRD0CIJNF&sprefix=serv%2Caps%2C355&ref=nb_sb_noss_2")

	// Create the CSV file
	file, err := os.Create("amazon_results.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the header row
	header := []string{"Title", "Price", "Image"}
	writer.Write(header)

	// Write the data rows
	for _, d := range data {
		writer.Write(d)
	}

	// Flush the writer
	writer.Flush()

	fmt.Println("Scraping finished. Data saved to amazon_results.csv")
}
