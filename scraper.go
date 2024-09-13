package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func joinArgs(args []string) string {
	var combinedStr string
	for i := 1; i < len(args); i++ {
		if i > 1 {
			combinedStr += " "
		}
		combinedStr += args[i]
	}

	return combinedStr
}

func main() {
	args := os.Args
	search := joinArgs(args)
	url := fmt.Sprintf("https://gogoanime3.co/search.html?keyword=%s", joinArgs(args))
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Responding from ", r.Request.URL)
	})

	collector.OnError((func(r *colly.Response, err error) { fmt.Println("fumble") }))

	collector.Visit(url)

}
