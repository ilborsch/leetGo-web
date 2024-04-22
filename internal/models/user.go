package models

import (
	"context"
)

type User struct {
	// uid is fetched from sso microservice therefore is not auto-incrementing
	UID            uint   `storage:"primaryKey;autoIncrement:false"`
	Username       string `storage:"not null;unique_index"`
	ProblemsSolved uint
	Articles       []Article
}

type UserProvider interface {
	User(ctx context.Context, uid uint) (User, error)
}

type UserAuthorizer interface {
	IsAdmin(ctx context.Context, uid uint) (bool, error)
}

type UserSaver interface {
	SaveUser(ctx context.Context, uid uint, username string) (uint, error)
}
