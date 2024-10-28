package user

import "github.com/google/uuid"

type UserDto struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	City     string    `json:"city"`
}
