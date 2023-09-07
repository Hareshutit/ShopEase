package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          *uuid.UUID
	UserID      *uuid.UUID
	Status      *bool
	Title       *string
	Description *string
	Price       *string
	Category    *string
	PathImages  *[]string
	Time        *time.Time
	Views       *int
}

type Parameters struct {
	Offset   *int
	Limit    *int
	Status   *bool
	Sort     *string
	UserId   *uuid.UUID
	Category *string
}
