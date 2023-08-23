package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config là struct chứa các giá trị cấu hình
type Config struct {
	PORT             string `mapstructure:"PORT"`
	DB_HOST          string `mapstructure:"DB_HOST"`
	DB_PORT          int    `mapstructure:"DB_PORT"`
	DB_USER          string `mapstructure:"DB_USER"`
	DB_PASSWORD      string `mapstructure:"DB_PASSWORD"`
	DB_NAME          string `mapstructure:"DB_NAME"`
	DB_POOL_MAX_OPEN int    `mapstructure:"DB_POOL_MAX_OPEN"`
	DB_POOL_MAX_IDLE int    `mapstructure:"DB_POOL_MAX_IDLE"`

	REDIS_ADDRESS           string `mapstructure:"REDIS_ADDRESS"`
	TIME_LIFE_CATCHING      int    `mapstructure:"TIME_LIFE_CATCHING"`
	REDIS_TIME_LIFE_CACHING int    `mapstructure:"REDIS_TIME_LIFE_CACHING"`
	REDIS_PORT              int    `mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD          string `mapstructure:"REDIS_PASSWORD"`

	PASSWORD_SALT string `mapstructure:"PASSWORD_SALT"`

	MAIL_USER string `mapstructure:"MAIL_USER"`
	MAIL_PASSWORD string `mapstructure:"MAIL_PASSWORD"`
	MAIL_HOST string `mapstructure:"MAIL_HOST"`
	MAIL_PORT string `mapstructure:"MAIL_PORT"`

	WEBSITE_URL string `mapstructure:"WEBSITE_URL"`
}

var Env Config

// LoadConfig là hàm dùng để load file cấu hình và lưu các giá trị vào Config
func InitConfig() {
	var config Config

	// Load file config
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	// Map các giá trị từ file config vào struct Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s ", err))
	}
	Env = config
}
