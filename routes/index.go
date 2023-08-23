package routes

import (
	"log"
	"net/http"
	"score_cat/apis"
	"score_cat/config"
	"score_cat/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	routerDefault := gin.Default()
	router := routerDefault.Group("/api")
	apis := []apis.IAPI{
		&apis.CrawJobAPI{},
	}
	router.Use(middlewares.ErrorMiddleware())
	for _, a := range apis {
		a.RegisterRoutes(router)
	}
	// Tạo router welcome
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my API!",
		})
	})

	// Xử lý lỗi 404
	routerDefault.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "API endpoint not found",
		})
	})

	err := routerDefault.Run(":" + config.Env.PORT)
	if err != nil {
		log.Fatal("Error starting up server:", err)
	}
}
