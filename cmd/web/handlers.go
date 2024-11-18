package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/ikhmal2/sapusapu/internal/db"
	"github.com/ikhmal2/sapusapu/internal/sqlQueries"
	"github.com/ikhmal2/sapusapu/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

var webUrl string = "https://gogoanime3.cc"

type animeItem struct {
	Name     string `json:"name"`
	Released string `json:"released"`
	Img      string `json:"img"`
	Link     string `json:"link"`
}

func getAnimeList(ctx *gin.Context) {
	keyword := ctx.Query("search")
	search := url.QueryEscape(keyword)
	url := fmt.Sprintf("%s/search.html?keyword=%s", webUrl, search)
	collector := colly.NewCollector()

	collector.OnError((func(r *colly.Response, err error) { fmt.Println("boy u fumbled:", err) }))
	collector.OnRequest((func(r *colly.Request) { fmt.Printf("Requesting to: %v\n", url) }))
	collector.OnResponse((func(r *colly.Response) { fmt.Printf("Got resp from %v", url) }))

	context := context.Background()
	var animeList []animeItem
	collector.OnHTML(".items", func(e *colly.HTMLElement) {
		anime := animeItem{}
		e.ForEach("li", func(i int, eachAnime *colly.HTMLElement) {
			anime.Name = eachAnime.ChildText("p.name")
			anime.Released = strings.Split(eachAnime.ChildText("p.released"), " ")[1]
			anime.Link = eachAnime.ChildAttr("a", "href")
			anime.Img = eachAnime.ChildAttr("img", "src")

			animeExits, animeData := utils.CheckExistingList(anime.Name)

			if !animeExits {
				_, err := db.DBconnect().InsertAnimeIntoList(context, sqlQueries.InsertAnimeIntoListParams{
					AnimeName: anime.Name,
					Released:  anime.Released,
					Img:       sql.NullString{String: anime.Img, Valid: true},
					Link:      anime.Link,
				})
				if err != nil {
					log.Println("Error executing query: ", err)
				}
				animeList = append(animeList, anime)
			} else {
				anime.Name = animeData.AnimeName
				anime.Released = animeData.Released.(string)
				anime.Link = animeData.Link
				anime.Img = animeData.Img.String
				animeList = append(animeList, anime)
			}

		})

	})

	if err := collector.Visit(url); err != nil {
		log.Println("error searching anime: ", err)
	}
	if len(animeList) != 0 {
		ctx.IndentedJSON(http.StatusOK, animeList)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, "Can't find the anime you're looking for")
	}
}

type animeEp struct {
	AnimeLink string `json:"link"`
	Episode   string `json:"episode"`
}

func goToAnime(ctx *gin.Context) {
	anime := ctx.Query("anime")
	animeReturned, err := utils.FindAnimeByLink(anime)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, "Can't find the anime you're looking for")
	}
	var animeEps []animeEp
	animePage := fmt.Sprintf("%s%s", webUrl, animeReturned.Link)

	collector := colly.NewCollector()
	collector.OnError((func(r *colly.Response, err error) { fmt.Println("\nboy u fumbled:", err) }))
	collector.OnRequest((func(r *colly.Request) { fmt.Printf("Requesting to: %v\n", animePage) }))

	// collector.OnHTML("#episode_related", func(e *colly.HTMLElement) {
	// 	fmt.Println("dah masuk")
	// 	animeEpisode := animeEp{}
	// 	e.ForEach("li", func(i int, eachEps *colly.HTMLElement) {
	// 		animeEpisode.AnimeLink = eachEps.ChildAttr("a", "href")
	// 		animeEpisode.Episode = eachEps.ChildText("div.name")
	// 		animeEps = append(animeEps, animeEpisode)
	// 	})
	// })
	collector.OnHTML("#load_ep", func(e *colly.HTMLElement) {
		fmt.Println("dah masuk")
		e.
	})

	if err := collector.Visit(animePage); err != nil {
		fmt.Println("Err while trying to visit: ", err)
	}
	if len(animeEps) > 0 {
		ctx.IndentedJSON(http.StatusOK, animeEps)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, "Can't find episodes for the anime")
	}
}
