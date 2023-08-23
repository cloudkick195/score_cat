package main

import (
	"score_cat/config"
	"score_cat/logger"
	"score_cat/repositories"
	"score_cat/routes"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			// Xử lý lỗi và ghi log
			logger.Error(err)
		}
	}()
	config.InitConfig()

	logger.Init()

	repositories.InitDatabase()
	repositories.AutoMigrate()
	routes.InitRouter()
}
