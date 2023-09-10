package migration

import (
	"log"

	"github.com/FianGumilar/e-wallet/infrastructure/database"
	"github.com/FianGumilar/e-wallet/models/entitty"
)

func Migration() {
	err := database.DbPG.AutoMigrate(
		&entitty.User{},
	)
	if err != nil {
		log.Printf("failed to migrate: %s", err)
	}
}
