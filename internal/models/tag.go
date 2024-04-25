package models

import (
	"context"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string    `gorm:"not null;unique"`
	Articles []Article `gorm:"many2many:article_tags;association_foreignkey:ID;foreignkey:ID"`
	Problems []Problem `gorm:"many2many:problem_tags;association_foreignkey:ID;foreignkey:ID"`
}

// TagRaw used to initialize tag object which will be later used for create/update operations
func TagRaw(name string) Tag {
	return Tag{Name: name}
}

type TagProvider interface {
	Tag(ctx context.Context, id int) (Tag, error)
	TagsByNames(ctx context.Context, names []string) ([]Tag, error)
	ArticleTags(ctx context.Context, articleID uint) ([]Tag, error)
	ProblemTags(ctx context.Context, problemID uint) ([]Tag, error)
}

type TagSaver interface {
	SaveTag(ctx context.Context, new Tag) (uint, error)
}

type TagRemover interface {
	RemoveTag(ctx context.Context, id uint) error
}
