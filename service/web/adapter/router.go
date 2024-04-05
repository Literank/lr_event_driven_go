/*
Package adapter adapts to all kinds of framework or protocols.
*/
package adapter

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"literank.com/event-books/domain/model"
	"literank.com/event-books/service/web/application"
	"literank.com/event-books/service/web/application/executor"
	"literank.com/event-books/service/web/infrastructure/config"
)

const (
	fieldOffset = "o"
	fieldQuery  = "q"
	fieldUID    = "uid"
)

// RestHandler handles all restful requests
type RestHandler struct {
	remoteConfig *config.RemoteServiceConfig
	bookOperator *executor.BookOperator
}

func newRestHandler(remote *config.RemoteServiceConfig, wireHelper *application.WireHelper) *RestHandler {
	return &RestHandler{
		remoteConfig: remote,
		bookOperator: executor.NewBookOperator(wireHelper.BookManager(), wireHelper.MessageQueueHelper()),
	}
}

// MakeRouter makes the main router
func MakeRouter(templates_pattern string, remote *config.RemoteServiceConfig, wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := newRestHandler(remote, wireHelper)
	// Create a new Gin router
	r := gin.Default()
	// Load HTML templates from the templates directory
	r.LoadHTMLGlob(templates_pattern)

	// Define a health endpoint handler
	r.GET("/", rest.indexPage)

	apiGroup := r.Group("/api")
	apiGroup.GET("/books", rest.getBooks)
	apiGroup.POST("/books", rest.createBook)
	return r, nil
}

// Render and show the index page
func (r *RestHandler) indexPage(c *gin.Context) {
	userID, err := c.Cookie(fieldUID)
	if err != nil {
		// Doesn't exist, make a new one
		userID = randomString(5)
		c.SetCookie(fieldUID, userID, 3600*24*30, "/", "", false, false)
	}
	q := c.Query(fieldQuery)
	books, err := r.bookOperator.GetBooks(c, 0, userID, q)
	if err != nil {
		c.String(http.StatusNotFound, "failed to get books")
		return
	}
	trends, err := r.bookOperator.GetTrends(c, r.remoteConfig.TrendURL)
	if err != nil {
		// It's not a must-have, just log the error
		log.Printf("Failed to get trends: %v", err)
		trends = make([]*model.Trend, 0)
	}
	// Render the HTML template named "index.html"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "LiteRank Book Store",
		"books":  books,
		"trends": trends,
		"q":      q,
	})
}

// Get all books
func (r *RestHandler) getBooks(c *gin.Context) {
	offset := 0
	offsetParam := c.Query(fieldOffset)
	if offsetParam != "" {
		value, err := strconv.Atoi(offsetParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		offset = value
	}
	books, err := r.bookOperator.GetBooks(c, offset, "", c.Query(fieldQuery))
	if err != nil {
		fmt.Printf("Failed to get books: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Create a new book
func (r *RestHandler) createBook(c *gin.Context) {
	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := r.bookOperator.CreateBook(c, &reqBody)
	if err != nil {
		fmt.Printf("Failed to create: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to create"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func randomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}
