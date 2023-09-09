package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FianGumilar/e-wallet/config"
	_ "github.com/lib/pq"
)

var DbPG *sql.DB

func GetDbPgConnection(conf *config.AppConfig) *sql.DB {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Name,
	)

	DbPG, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("failed connecting to database: %v", err)
	}

	err = DbPG.Ping()
	if err != nil {
		log.Printf("failed connecting to database: %v", err)
	}

	return DbPG
}
