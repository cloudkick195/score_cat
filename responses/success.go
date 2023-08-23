package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": message,
	})
}
