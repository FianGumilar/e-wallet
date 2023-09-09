package repository

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/e-wallet/interfaces"
	"github.com/FianGumilar/e-wallet/models/entitty"
)

type repository struct {
	db *sql.DB
}

func NewUserRepository(con *sql.DB) interfaces.UserRepository {
	return &repository{db: con}
}

// FindByID implements interfaces.UserRepository.
func (r repository) FindByID(ctx context.Context, id int64) (user entitty.User, err error) {
	query := `SELECT * FROM users WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone)
	if err != nil {
		return user, nil
	}
	return
}

// FindByUsername implements interfaces.UserRepository.
func (r repository) FindByUsername(ctx context.Context, username string) (user entitty.User, err error) {
	query := `SELECT * FROM users WHERE username = ?`

	row := r.db.QueryRowContext(ctx, query, username)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone)
	if err != nil {
		return user, nil
	}
	return
}
