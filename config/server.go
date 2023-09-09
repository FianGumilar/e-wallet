package config

import "os"

func iniServer(conf *AppConfig) {
	port := os.Getenv("PORT")

	conf.Server.Port = port
}
