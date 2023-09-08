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

func (s *Post) New() {
	s.Id = new(uuid.UUID)
	s.UserID = new(uuid.UUID)
	s.Status = new(bool)
	s.Title = new(string)
	s.Description = new(string)
	s.Price = new(string)
	s.Category = new(string)
	s.PathImages = new([]string)
	s.Time = new(time.Time)
	s.Views = new(int)
}
