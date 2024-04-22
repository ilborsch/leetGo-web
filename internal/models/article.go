package models

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title       string
	Content     []byte `gorm:"type:blob"`
	AuthorID    uint
	IsPublished bool
	PublishDate time.Time
	Tags        []Tag `gorm:"many2many:article_tags;"`
}

type ArticleProvider interface {
	Article(ctx context.Context, id uint) (*Article, error)
	ArticlesByAuthor(ctx context.Context, authorID uint) ([]Article, error)
	ArticlesByTags(ctx context.Context, tags []Tag) ([]Article, error)
}

type ArticleSaver interface {
	Save(
		ctx context.Context,
		id uint,
		title string,
		content []byte,
		authorID uint,
		isPublished bool,
		tags []Tag,
	) (uint, error)
}

type ArticleUpdater interface {
	Update(ctx context.Context, id uint, new Article) error
}

type ArticleRemover interface {
	Remove(ctx context.Context, id uint) error
}
