package models

import (
	"WeatherfForecast/dto/user"
)

type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	City     string `gorm:"not null" json:"city"`
}

func (u User) ToUserDto() user.UserDto {
	return user.UserDto{
		Id:       u.Id,
		Username: u.Username,
		City:     u.City,
	}
}
