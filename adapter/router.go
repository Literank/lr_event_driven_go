/*
Package adapter adapts to all kinds of framework or protocols.
*/
package adapter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"literank.com/event-books/application"
	"literank.com/event-books/application/executor"
	"literank.com/event-books/domain/model"
)

const (
	fieldOffset = "o"
	fieldQuery  = "q"
)

// RestHandler handles all restful requests
type RestHandler struct {
	bookOperator *executor.BookOperator
}

func newRestHandler(wireHelper *application.WireHelper) *RestHandler {
	return &RestHandler{
		bookOperator: executor.NewBookOperator(wireHelper.BookManager()),
	}
}

// MakeRouter makes the main router
func MakeRouter(templates_pattern string, wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := newRestHandler(wireHelper)
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
	q := c.Query(fieldQuery)
	books, err := r.bookOperator.GetBooks(c, 0, q)
	if err != nil {
		c.String(http.StatusNotFound, "failed to get books")
		return
	}
	// Render the HTML template named "index.html"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "LiteRank Book Store",
		"books": books,
		"q":     q,
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
	books, err := r.bookOperator.GetBooks(c, offset, c.Query(fieldQuery))
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
