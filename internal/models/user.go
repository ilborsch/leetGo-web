package models

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"not null;unique_index"`
	ProblemsSolved uint
	Articles       []Article `gorm:"foreignkey:AuthorID"`
}

type UserProvider interface {
	User(ctx context.Context, uid uint) (User, error)
}

type UserSaver interface {
	SaveUser(ctx context.Context, uid uint, username string) (uint, error)
}
