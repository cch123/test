package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

var visited = map[string]bool{}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.v2ex.com"),
		colly.MaxDepth(1),
	)

	detailRegex, _ := regexp.Compile(`/go/go\?p=\d+$`)
	listRegex, _ := regexp.Compile(`/t/\d+#\w+`)
	//panRegex, _ := regexp.Compile(`pan.baidu.com`)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		// Print article
		/*
			if !panRegex.Match([]byte(e.Text)) {
				return
			}
		*/

		fmt.Println(e.Text)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// 已访问过的详情页或列表页，跳过
		if visited[link] && (detailRegex.Match([]byte(link)) || listRegex.Match([]byte(link))) {
			return
		}

		// 匹配下列两种 url 模式的，才去 visit
		// https://www.v2ex.com/go/go?p=2
		// https://www.v2ex.com/t/472945#reply3
		if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
			println("not match", link)
			return
		}
		time.Sleep(time.Second)
		println("match", link)

		visited[link] = true

		time.Sleep(time.Millisecond * 2)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		/*
			r.Headers.Set("Cookie", "")
			r.Headers.Set("DNT", "1")
			r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
			r.Headers.Set("Host", "www.v2ex.com")
		*/
	})

	err := c.Visit("https://www.v2ex.com/go/go")
	if err != nil {
		fmt.Println(err)
	}
}
