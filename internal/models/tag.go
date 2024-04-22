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

// TagRaw used to initialize tag object which will be later used for create/update operations
func TagRaw(name string) Tag {
	return Tag{Name: name}
}

type TagProvider interface {
	Tag(ctx context.Context, id int) (Tag, error)
	TagByName(ctx context.Context, name string) (Tag, error)
}

type TagSaver interface {
	SaveTag(ctx context.Context, new Tag) (uint, error)
}

type TagRemover interface {
	RemoveTag(ctx context.Context, id uint) error
}
