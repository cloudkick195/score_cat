package apis

import (
	"net/http"
	"score_cat/responses"
	"score_cat/services"

	"github.com/gin-gonic/gin"
)

type CrawJobAPI struct{}

func (a *CrawJobAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.Group("/crawjob").
		GET("/", func(c *gin.Context) {
			crawService := services.CrawService{}

			err := crawService.CrawlAndSaveData()
			if err != nil {
				responses.ResponseErr(c, http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Sucess"})
		})
}
