package middlewares

import (
	"net/http"
	"score_cat/logger"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Xử lý lỗi và ghi log
				logger.Log.Error("Panic occurred", err.(error))

				// Trả về response lỗi cho client
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Xử lý lỗi và ghi log
			logger.Log.Error("Error occurred", err)

			// Trả về response lỗi cho client
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
