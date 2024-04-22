package storage

import (
	"context"
	"github.com/ilborsch/leetGo-web/internal/models"
	"github.com/ilborsch/leetGo-web/internal/utils"
)

func (s *Storage) Article(ctx context.Context, id uint) (models.Article, error) {
	var article models.Article
	if res := s.db.First(&article, id); res.Error != nil {
		return models.Article{}, res.Error
	}
	return article, nil
}

func (s *Storage) ArticlesByAuthor(ctx context.Context, authorID uint) ([]models.Article, error) {
	var articles []models.Article
	result := s.db.Where("author_id = ?", authorID).Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (s *Storage) ArticlesByTags(ctx context.Context, tags []models.Tag) ([]models.Article, error) {
	var articles []models.Article
	tagNames := utils.GetTagNames(tags)
	// Query using JOIN to filter articles by multiple tags
	result := s.db.Select("id, title, content, author_id, is_published, publish_date").
		Joins("JOIN article_tags ON articles.id = article_tags.article_id").
		Joins("JOIN tags ON tags.id = article_tags.tag_id").
		Where("tags.name IN ?", tagNames).
		Group("articles.id").
		Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (s *Storage) UpdateArticle(ctx context.Context, id uint, new models.Article) error {
	var article models.Article
	result := s.db.First(&article, id)
	if result.Error != nil {
		return result.Error
	}

	result = s.db.Model(&article).Updates(new)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) RemoveArticle(ctx context.Context, id uint) error {
	result := s.db.Delete(&models.Article{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
