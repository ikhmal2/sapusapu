package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

var webUrl string = "https://gogoanime3.co/"

type animeItem struct {
	Name     string `json:"name"`
	Released string `json:"released"`
	Img      string `json:"img"`
	Link     string `json:"link"`
}

func getAnimeList(ctx *gin.Context) {
	keyword := ctx.Query("search")
	search := url.QueryEscape(keyword)
	url := fmt.Sprintf("%ssearch.html?keyword=%s", webUrl, search)
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
		// fmt.Printf("anime list %v", animeList)
	})

	collector.Visit(url)
	ctx.IndentedJSON(http.StatusOK, animeList)
}

func main() {
	router := gin.Default()
	router.POST("/getanime", getAnimeList)
	router.Run("localhost:8080")
}
