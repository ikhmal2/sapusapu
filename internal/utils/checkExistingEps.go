package utils

import (
	"context"
	"log"

	"github.com/ikhmal2/sapusapu/internal/db"
	"github.com/ikhmal2/sapusapu/internal/sqlQueries"
)

func CheckExistingEp(episodeData *sqlQueries.GetAnimeEpisodeParams) bool {
	context := context.Background()
	_, err := db.DBconnect().GetAnimeEpisode(context, *episodeData)
	if err != nil {
		log.Println("Error: ", err)
		return false
	} else {
		return true
	}
}
