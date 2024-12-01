package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
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

func getAnimeEps(contx *gin.Context) {
	anime := contx.Query("anime")
	animeReturned, err := utils.FindAnimeByLink(anime)
	animeID := animeReturned.AnimeID
	if err != nil {
		contx.IndentedJSON(http.StatusInternalServerError, "Can't find the anime you're looking for")
	}
	url := fmt.Sprintf("%s%s", webUrl, animeReturned.Link)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2000*time.Millisecond),
		chromedp.ActionFunc(func(ctx context.Context) error {
			rootNode, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			html, err = dom.GetOuterHTML().WithNodeID(rootNode.NodeID).Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("Err trying to parse HTML:", err)
	}

	context := context.Background()

	episodeRelated := doc.Find("#episode_related > li > a")
	var episodes []animeEp
	episodeItem := animeEp{}
	var episodeCheck *sqlQueries.GetAnimeEpisodeParams
	episodeRelated.Each(func(_ int, s *goquery.Selection) {
		episodeItem.AnimeLink, _ = s.Attr("href")
		episodeItem.Episode = s.Children().Text()
		episodeCheck.Animeid = sql.NullInt64{Int64: animeID, Valid: true}
		episodeCheck.Episode = episodeItem.AnimeLink

		if !utils.CheckExistingEp(episodeCheck) {
			_, err := db.DBconnect().InsertAnimeEp(context, sqlQueries.InsertAnimeEpParams{
				Animeid: sql.NullInt64{Int64: animeID, Valid: true},
				Episode: episodeItem.AnimeLink,
			})
			if err != nil {
				log.Println("Error inserting episode: ", err)
			}
			episodes = append(episodes, episodeItem)
		} else {
			episodes = append(episodes, episodeItem)
		}

	})

	if len(episodes) != 0 {
		contx.IndentedJSON(http.StatusOK, episodes)
	} else {
		contx.IndentedJSON(http.StatusNotFound, "Can't find the episodes you're looking for")
	}
}
