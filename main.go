package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type animeItem struct {
	Name     string `json:"name"`
	Released string `json:"released"`
	Img      string `json:"img"`
	Link     string `json:"link"`
}

func joinArgs(args string) string {
	var combinedStr string
	tempStr := strings.Split(args, " ")
	for i := 1; i < len(tempStr); i++ {
		if i > 1 {
			combinedStr += " "
		}
		combinedStr += tempStr[i]
	}
	return combinedStr
}

func getAnimeList(ctx *gin.Context) {
	keyword := ctx.Param("keyword")
	// search := url.QueryEscape(joinArgs(keyword))
	search := url.QueryEscape(keyword)
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
		fmt.Printf("anime list %v", animeList)
	})

	collector.Visit(url)
	ctx.IndentedJSON(http.StatusOK, animeList)
	// return animeList
}

func main() {
	router := gin.Default()
	router.POST("/getanime", getAnimeList)

	router.Run("localhost:8080")
}
