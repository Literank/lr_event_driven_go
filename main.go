package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Load HTML templates from the templates directory
	router.LoadHTMLGlob("templates/*.html")

	// Define a route for the homepage
	router.GET("/", func(c *gin.Context) {
		// Render the HTML template named "index.html"
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "LiteRank Book Store",
		})
	})

	// Run the server, default port is 8080
	router.Run()
}
