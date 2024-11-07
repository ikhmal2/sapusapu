package utils

import (
	"context"
	"log"

	"github.com/ikhmal2/sapusapu/internal/db"
	"github.com/ikhmal2/sapusapu/internal/sqlQueries"
)

func CheckExistingList(aniName string) (bool, *sqlQueries.AnimeList) {
	context := context.Background()
	anime, err := db.DBconnect().FindAnime(context, aniName)
	if err != nil {
		log.Println("Error:", err)
		return false, nil
	} else {
		return true, &anime
	}
}
