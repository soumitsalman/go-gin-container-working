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

func main() {
	log.Println("Starting app")
	router := gin.Default()

	router.GET("/ping", handlePing)

	router.Run()
}
