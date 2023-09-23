package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Server is starting...")
}

func defaultRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}