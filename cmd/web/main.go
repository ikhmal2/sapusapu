package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/getanime", getAnimeList)
	router.Run("localhost:8080")
}
