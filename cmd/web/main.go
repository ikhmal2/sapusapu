package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/anime/v1/get-anime", getAnimeList)
	router.POST("/anime/v1/get-episodes", getAnimeEps)
	router.Run("localhost:8080")
}
