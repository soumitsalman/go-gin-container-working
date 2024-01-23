package main

import (
	"log"
	"net/http"

	"examples.com/gowebapp/shared"
	"github.com/gin-gonic/gin"
)

func handlePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, shared.GetPong())
}

func handleContents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, shared.GetContents())
}

func main() {
	log.Println("Starting app")
	router := gin.Default()

	router.GET("/ping", handlePing)
	router.GET("/contents", handleContents)

	router.Run()
}
