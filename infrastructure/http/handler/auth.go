package handler

import (
	"log"
	"net/http"

	middlewares "github.com/FianGumilar/e-wallet/infrastructure/http/middlewares"
	"github.com/FianGumilar/e-wallet/interfaces"
	"github.com/FianGumilar/e-wallet/models/dto"
	"github.com/labstack/echo/v4"
)

type handler struct {
	userService interfaces.UserService
}

func NewAuthHandler(e *echo.Echo, userService interfaces.UserService) {
	h := &handler{
		userService: userService,
	}
	e.POST("token/generate", h.GenerateToken)
	e.GET("token/validate", h.ValidateToken, middlewares.Authenticate(userService))
}

// GenerateToken implements interfaces.UserHandler.
func (h handler) GenerateToken(c echo.Context) error {
	var req dto.AuthReq

	// Pasrsing Body Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to parse request body")
	}

	// authenticate token with user service
	token, err := h.userService.Authenticate(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Failed to authenticate")
	}
	log.Printf("Generated token: %s", token.Token)

	return c.JSON(http.StatusOK, token)
}

// ValidateToken implements interfaces.UserHandler.
func (h handler) ValidateToken(c echo.Context) error {
	user := c.Get("x-user")

	return c.JSON(http.StatusOK, user)
}
