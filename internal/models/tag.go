package models

import (
	"context"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string    `storage:"not null;unique;"`
	Articles []Article `storage:"many2many:article_tags;"`
	Problems []Problem `storage:"many2many:problem_tags;"`
}

type TagProvider interface {
	Tag(ctx context.Context, name string) (Tag, error)
}

type TagSaver interface {
	SaveTag(ctx context.Context, name string) (uint, error)
}

type TagRemover interface {
	RemoveTag(ctx context.Context, id uint) error
}
