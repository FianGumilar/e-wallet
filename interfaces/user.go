package interfaces

import (
	"context"

	"github.com/FianGumilar/e-wallet/models/dto"
	"github.com/FianGumilar/e-wallet/models/entitty"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (entitty.User, error)
	FindByUsername(ctx context.Context, username string) (entitty.User, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (res dto.AuthRes, err error)
	Validate(ctx context.Context, token string) (user dto.UserData, err error)
}

type UserHandler interface {
	GenerateToken(c echo.Context) error
	ValidateToken(c echo.Context) error
}
