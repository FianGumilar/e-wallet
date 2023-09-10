package database

import (
	"fmt"
	"log"

	"github.com/FianGumilar/e-wallet/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbPG *gorm.DB

func GetDbPgConnection(conf *config.AppConfig) *gorm.DB {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Name,
	)

	DbPG, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed connecting to database: %v", err)
	}

	return DbPG
}
