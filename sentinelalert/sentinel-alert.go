package sentinelalert

import (
	"net/http"
	"az-function/models"
	"github.com/gin-gonic/gin"
)

// HandleSentinel handle sentinel POST request
func HandleSentinel(c *gin.Context) {
	var alert models.SentinelAlert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "body received"})
}
