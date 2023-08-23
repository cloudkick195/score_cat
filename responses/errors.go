package responses

import (
	"net/http"
	"score_cat/types"

	"github.com/gin-gonic/gin"
)

func ResponsePanic(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func ResponseErr(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}

func CheckErrors(results ...types.Result) error {
	for _, result := range results {
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
