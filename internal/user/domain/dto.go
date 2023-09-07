package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID

	Email       string
	PhoneNumber string
	Address     string

	Login         string
	Password      string
	CheckPassword string

	Name string

	PathToAvatar string
}
