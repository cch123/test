package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var domain2Collector = map[string]*colly.Collector{}

func factory(domain string) *colly.Collector {
	return domain2Collector[domain]
}

func init() {
	v2exCollector := colly.NewCollector(
		colly.AllowedDomains("www.v2ex.com"),
		colly.MaxDepth(0),
	)

	v2exCollector.OnResponse(func(resp *colly.Response) {
		//fmt.Println(string(resp.Body))
	})

	v2exCollector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})
	domain2Collector["https://www.v2ex.com/go/go"] = v2exCollector

	daydownCollector := colly.NewCollector(
		colly.AllowedDomains("www.v2ex.com"),
		colly.MaxDepth(0),
	)

	daydownCollector.OnResponse(func(resp *colly.Response) {
		//fmt.Println(string(resp.Body))
	})

	daydownCollector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})
	domain2Collector["https://0daydown.com"] = daydownCollector
}

func main() {
	urls := []string{"https://www.v2ex.com/go/go", "https://0daydown.com"}
	for _, url := range urls {
		instance := factory(url)
		instance.Visit(url)
	}
}
