package repository

import (
	"context"

	"github.com/FianGumilar/e-wallet/interfaces"
	"github.com/FianGumilar/e-wallet/models/entitty"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(con *gorm.DB) interfaces.UserRepository {
	return &repository{db: con}
}

// FindByID implements interfaces.UserRepository.
func (r repository) FindByID(ctx context.Context, id int64) (user entitty.User, err error) {
	dataset := r.db.Where("id = ?", id).First(&user)
	if dataset.Error != nil {
		return user, nil
	}
	return
}

// FindByUsername implements interfaces.UserRepository.
func (r repository) FindByUsername(ctx context.Context, username string) (user entitty.User, err error) {
	dataset := r.db.Where("username = ?", username).First(&user)
	if dataset.Error != nil {
		return user, nil
	}
	return
}
