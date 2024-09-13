package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type animeItem struct {
	Name     string
	Released string
	Img      string
	Link     string
}

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
	search := url.QueryEscape(joinArgs(args))
	url := fmt.Sprintf("https://gogoanime3.co/search.html?keyword=%s", search)
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Responding from ", r.Request.URL)
	})

	collector.OnError((func(r *colly.Response, err error) { fmt.Println("boy u fumbled") }))

	var animeList []animeItem
	collector.OnHTML(".items", func(e *colly.HTMLElement) {
		anime := animeItem{}
		e.ForEach("li", func(i int, eachAnime *colly.HTMLElement) {
			anime.Name = eachAnime.ChildText("p.name")
			anime.Released = eachAnime.ChildText("p.released")
			anime.Link = eachAnime.ChildAttr("a", "href")
			anime.Img = eachAnime.ChildAttr("img", "src")
			animeList = append(animeList, anime)
		})
		fmt.Println(animeList)
	})

	collector.Visit(url)

}
