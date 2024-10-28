package repository

import (
	"WeatherfForecast/dto/user"
	"WeatherfForecast/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	ChangeCity(username string, request user.ChangeCityRequest) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) Create(user *models.User) error {
	return u.db.Create(user).Error
}

func (u userRepository) FindByUsername(username string) (*models.User, error) {
	var foundUser models.User

	if err := u.db.Where("username = ?", username).First(&foundUser).Error; err != nil {
		return nil, err
	}
	return &foundUser, nil
}

func (u userRepository) ChangeCity(username string, request user.ChangeCityRequest) (models.User, error) {
	foundUser, err := u.FindByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	foundUser.City = request.City

	if err := u.db.Save(foundUser).Error; err != nil {
		return models.User{}, err
	}

	return *foundUser, nil
}
