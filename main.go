package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ankitdmon/e-commerce-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router := gin.New()
	router.Use((gin.Logger()))
	routes.SetupRoutes(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Golang",
		})
	})
	log.Fatal(router.Run(":" + PORT))
}
