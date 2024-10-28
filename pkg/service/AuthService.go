package service

import (
	"WeatherfForecast/dto/user"
	"WeatherfForecast/models"
	"WeatherfForecast/pkg/repository"
	"WeatherfForecast/util"
	"errors"
	"fmt"
)

type AuthService interface {
	Register(userRegisterRequest user.RegisterUserRequest) (*models.User, error)
	Login(userLoginRequest user.LoginUserRequest) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
	jwtUtil        util.JwtUtil
}

func NewAuthService(userRepository repository.UserRepository, jwtUtil util.JwtUtil) AuthService {
	return &authService{
		userRepository: userRepository,
		jwtUtil:        jwtUtil,
	}
}

func (a authService) Register(userRegisterRequest user.RegisterUserRequest) (*models.User, error) {
	existingUser, _ := a.userRepository.FindByUsername(userRegisterRequest.Username)

	if existingUser != nil {
		return nil, errors.New(fmt.Sprintf("User with username %s already exists", userRegisterRequest.Username))
	}

	hashedPassword, err := util.HashPassword(userRegisterRequest.Password)

	if err != nil {
		return nil, err
	}

	createdUser := &models.User{
		Username: userRegisterRequest.Username,
		Password: hashedPassword,
		City:     userRegisterRequest.City,
	}

	if err := a.userRepository.Create(createdUser); err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (a authService) Login(userLoginRequest user.LoginUserRequest) (string, error) {
	dbUser, err := a.userRepository.FindByUsername(userLoginRequest.Username)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !util.CheckPasswordHash(userLoginRequest.Password, dbUser.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := a.jwtUtil.GenerateToken(dbUser.Id, dbUser.Username)

	if err != nil {
		return "", err
	}

	return token, nil
}
