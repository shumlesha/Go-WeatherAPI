package service

import (
	"WeatherfForecast/dto/user"
	"WeatherfForecast/models"
	"WeatherfForecast/pkg/repository"
)

type UserService interface {
	ChangeCity(username string, changeCityRequest user.ChangeCityRequest) (models.User, error)
	GetMyProfile(username string) (*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) GetMyProfile(username string) (*models.User, error) {
	return u.userRepository.FindByUsername(username)
}

func (u userService) ChangeCity(username string, changeCityRequest user.ChangeCityRequest) (models.User, error) {
	return u.userRepository.ChangeCity(username, changeCityRequest)
}
