/*
Package adapter adapts to all kinds of framework or protocols.
*/
package adapter

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"literank.com/event-books/service/trend/application"
	"literank.com/event-books/service/trend/application/executor"
)

const (
	fieldOffset = "o"
)

// RestHandler handles all restful requests
type RestHandler struct {
	trendOperator *executor.TrendOperator
}

func newRestHandler(wireHelper *application.WireHelper) *RestHandler {
	return &RestHandler{
		trendOperator: executor.NewTrendOperator(wireHelper.TrendManager()),
	}
}

// MakeRouter makes the main router
func MakeRouter(wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := newRestHandler(wireHelper)
	// Create a new Gin router
	r := gin.Default()

	// Define a health endpoint handler
	r.GET("/trends", rest.getTrends)
	return r, nil
}

// Get all trends
func (r *RestHandler) getTrends(c *gin.Context) {
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
	trends, err := r.trendOperator.TopTrends(c, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get trends"})
		return
	}
	c.JSON(http.StatusOK, trends)
}
