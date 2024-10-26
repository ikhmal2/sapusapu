package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/ikhmal2/sapusapu/internal/db"
	"github.com/ikhmal2/sapusapu/internal/sqlQueries"
	_ "github.com/mattn/go-sqlite3"
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

	collector.OnError((func(r *colly.Response, err error) { fmt.Println("boy u fumbled:", err) }))

	var animeList []animeItem
	collector.OnHTML(".items", func(e *colly.HTMLElement) {
		anime := animeItem{}
		e.ForEach("li", func(i int, eachAnime *colly.HTMLElement) {
			anime.Name = eachAnime.ChildText("p.name")
			anime.Released = eachAnime.ChildText("p.released")
			anime.Link = eachAnime.ChildAttr("a", "href")
			anime.Img = eachAnime.ChildAttr("img", "src")
			animeList = append(animeList, anime)

			insertedAnime, err := db.DBconnect().InsertAnimeIntoList(ctx, sqlQueries.InsertAnimeIntoListParams{
				AnimeName: anime.Name,
				Released:  anime.Released,
				Img:       sql.NullString{String: anime.Img, Valid: true},
				Link:      anime.Link,
			})

			if err != nil {
				log.Print("Error executing query: ", err)
			}
			log.Print("Anime inserted into DB: ", insertedAnime)
		})

	})

	collector.Visit(url)
	ctx.IndentedJSON(http.StatusOK, animeList)
}
