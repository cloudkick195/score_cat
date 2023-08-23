package repositories

import (
	"fmt"
	"log"
	"os"
	"score_cat/config"
	"score_cat/models"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Ghi log ra stdout
		logger.Config{
			SlowThreshold: time.Second, // Thời gian mà các truy vấn được xem như chậm (nếu thực thi lâu hơn, sẽ được ghi lại trong log)
			LogLevel:      logger.Info, // Mức độ log
			Colorful:      true,        // Sử dụng màu sắc cho log
		},
	)
	dbURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Env.DB_USER,
		config.Env.DB_PASSWORD,
		config.Env.DB_HOST,
		strconv.Itoa(config.Env.DB_PORT),
		config.Env.DB_NAME,
	)

	conn, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	DB = conn
}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.Team{},
		&models.Country{},
		&models.League{},
	)
	//insertLargeData()
}
