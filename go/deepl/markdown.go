package main

import (
	"fmt"
	"io/ioutil"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func main() {
	markdown := goldmark.New(
		goldmark.WithExtensions(),
	)

	contentBytes, err := ioutil.ReadFile("./md.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	doc := markdown.Parser().Parse(text.NewReader(contentBytes))
	fmt.Println(doc)
	fmt.Println(doc.ChildCount())
	c := doc.FirstChild()
	for c != nil {
		c.Dump(contentBytes, 0)
		c = c.NextSibling()
	}
}
