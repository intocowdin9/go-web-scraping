package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
	Cover string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com", "www.books.toscrape.com"),
	)

	scrapeUrl := "http://books.toscrape.com/"

	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	hearders := []string{"Title", "Price", "Cover"}
	writer.Write(hearders)

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		book.Cover = e.ChildAttr("img.thumbnail", "src")

		row := []string{book.Title, book.Price, book.Cover}
		fmt.Println(row)
		writer.Write(row)

	})

	c.OnHTML(".next  a", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(scrapeUrl)
}
