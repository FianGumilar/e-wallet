package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/FianGumilar/e-wallet/interfaces"
	"github.com/FianGumilar/e-wallet/models/dto"
	"github.com/FianGumilar/e-wallet/models/entitty"
	"github.com/FianGumilar/e-wallet/utils"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository  interfaces.UserRepository
	cacheRepository interfaces.CacheRepository
}

func NewUserService(userRepository interfaces.UserRepository, cacheRepository interfaces.CacheRepository) interfaces.UserService {
	return &service{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
	}
}

// Authenticate implements interfaces.UserService.
func (s service) Authenticate(ctx context.Context, req dto.AuthReq) (res dto.AuthRes, err error) {
	// Check user by FindByUsername
	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, utils.ErrAuthFailed
	}

	// Check user exists
	if user == (entitty.User{}) {
		return dto.AuthRes{}, err
	}

	// Compare hash & password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, err
	}

	// Generate auth token
	token := utils.GenerateRandomString(32)
	log.Printf("Generated Token: %s", token)

	// Set token as a cache
	userJson, _ := json.Marshal(user)
	_ = s.cacheRepository.Set("user:"+token, userJson)

	return dto.AuthRes{
		UserID: user.ID,
		Token:  token,
	}, nil

}

// Validate implements interfaces.UserService.
func (s service) Validate(ctx context.Context, token string) (user dto.UserData, err error) {
	data, err := s.cacheRepository.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, utils.ErrAuthFailed
	}

	var users entitty.User
	_ = json.Unmarshal(data, &users)

	return dto.UserData{
		ID:       user.ID,
		Username: user.Username,
		Phone:    user.Phone,
	}, nil
}
