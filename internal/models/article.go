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
	AuthorID    uint   `gorm:"column:author_id"`
	IsPublished bool
	PublishDate time.Time
	Tags        []Tag `gorm:"many2many:article_tags;association_foreignkey:ID;foreignkey:ID"`
}

// ArticleRaw used to initialize article object which will be later used for create/update operations
func ArticleRaw(title string, content []byte, authorID uint, isPublished bool, tags []Tag) Article {
	publishDate := time.Time{}
	if isPublished {
		publishDate = time.Now()
	}
	return Article{
		Title:       title,
		Content:     content,
		AuthorID:    authorID,
		IsPublished: isPublished,
		PublishDate: publishDate,
		Tags:        tags,
	}
}

type ArticleProvider interface {
	Article(ctx context.Context, id uint) (Article, error)
	ArticlesByAuthor(ctx context.Context, authorID uint) ([]Article, error)
	ArticlesByTags(ctx context.Context, tags []Tag) ([]Article, error)
}

type ArticleSaver interface {
	SaveArticle(ctx context.Context, new Article) (uint, error)
}

type ArticleUpdater interface {
	UpdateArticle(ctx context.Context, id uint, new Article) error
}

type ArticleRemover interface {
	RemoveArticle(ctx context.Context, id uint) error
}
