package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"regexp"

	"github.com/gocolly/colly"
)

var f *os.File

var visited = map[string]bool{}

var fileName = "./crawl_log"

// open file, read links already crawled
func init() {
	contents, err := ioutil.ReadFile(fileName)
	fmt.Println(string(contents), err)
	links := strings.Split(string(contents), "\n")

	for _, link := range links {
		if link != "" {
			visited[link] = true
		}
	}

	// fmt.Printf("%#v\n", visited)

	f, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var urlMap = map[string]bool{}

var separater = "#######################################################"

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.0daydown.com"),
	)
	re1, _ := regexp.Compile(`.*//www.0daydown.com/\d+/\d+.html$`)
	re2, _ := regexp.Compile(`.*//www.0daydown.com/category/tutorials/page/\d+`)
	re3, _ := regexp.Compile(`pan.baidu.com`)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		// Print article
		if !re3.Match([]byte(e.Text)) {
			return
		}

		fmt.Println(e.Text)
		fmt.Println(separater)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// 已访问过的详情页，跳过
		if urlMap[link] && re1.Match([]byte(link)) {
			return
		}

		// 匹配下列两种 url 模式的，才去 visit
		// http://www.0daydown.com/category/tutorials/page/6
		// http://www.0daydown.com/07/896908.html
		if !re1.Match([]byte(link)) && !re2.Match([]byte(link)) {
			return
		}

		urlMap[link] = true
		f.WriteString(link + "\n")

		time.Sleep(time.Millisecond * 100)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", "")
		r.Headers.Set("DNT", "1")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
		r.Headers.Set("Host", "www.0daydown.com")

		// fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://www.0daydown.com/category/tutorials/page/3129")
}
