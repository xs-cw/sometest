package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetList ...
func GetList(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{"time": "2019-12-01", "title": "first title", "info": "first info for card"},
		{"time": "2019-12-02", "title": "second title", "info": "second info for card"},
		{"time": "2019-12-03", "title": "third title", "info": "third info for card"},
	})
}
