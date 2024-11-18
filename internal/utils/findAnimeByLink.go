package utils

import (
	"context"
	"log"

	"github.com/ikhmal2/sapusapu/internal/db"
	"github.com/ikhmal2/sapusapu/internal/sqlQueries"
)

func FindAnimeByLink(link string) (*sqlQueries.AnimeList, error) {
	context := context.Background()
	anime, err := db.DBconnect().GetAnimeEpsByLink(context, link)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	} else {
		log.Println("Found anime by link", anime)
		return &anime, nil
	}
}
