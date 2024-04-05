/*
Package adapter adapts to all kinds of framework or protocols.
*/
package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"literank.com/event-books/service/recommendation/application"
	"literank.com/event-books/service/recommendation/application/executor"
)

const (
	fieldUID = "uid"
)

// RestHandler handles all restful requests
type RestHandler struct {
	interestOperator *executor.InterestOperator
}

func newRestHandler(wireHelper *application.WireHelper) *RestHandler {
	return &RestHandler{
		interestOperator: executor.NewInterestOperator(wireHelper.InterestManager()),
	}
}

// MakeRouter makes the main router
func MakeRouter(wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := newRestHandler(wireHelper)
	// Create a new Gin router
	r := gin.Default()

	// Define a health endpoint handler
	r.GET("/recommendations", rest.getInterests)
	return r, nil
}

// Get all trends
func (r *RestHandler) getInterests(c *gin.Context) {
	uid := c.Query(fieldUID)
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty uid is not allowed"})
		return
	}
	trends, err := r.interestOperator.InterestsForUser(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get interests for user " + uid})
		return
	}
	c.JSON(http.StatusOK, trends)
}
