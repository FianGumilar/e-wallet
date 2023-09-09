package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Env string
	}
	Server struct {
		Port string
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
		Ssl  string
	}
	Redis struct {
		Addr string
		Pass string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error load .env file %v:", err)
	}
	if appConfig == nil {
		appConfig = &AppConfig{}
		initApp(appConfig)
		initDB(appConfig)
		initRedis(appConfig)
	}

	return appConfig

}
