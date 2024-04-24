package storage

import (
	"context"
	"errors"
	"github.com/ilborsch/leetGo-web/internal/models"
	"gorm.io/gorm"
)

func (s *Storage) Tag(ctx context.Context, id int) (models.Tag, error) {
	var tag models.Tag
	result := s.db.First(&tag, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Tag{}, nil
		}
		return models.Tag{}, result.Error
	}
	return tag, nil
}

func (s *Storage) TagsByNames(ctx context.Context, names []string) ([]models.Tag, error) {
	var tags []models.Tag
	result := s.db.Where("name IN ?", names).Find(&tags)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return tags, nil
}

func (s *Storage) ArticleTags(ctx context.Context, articleID uint) ([]models.Tag, error) {
	var article models.Article
	if result := s.db.Preload("Tags").First(&article, articleID); result.Error != nil {
		return nil, result.Error
	}
	return article.Tags, nil
}

func (s *Storage) ProblemTags(ctx context.Context, problemID uint) ([]models.Tag, error) {
	var problem models.Problem
	if result := s.db.Preload("Tags").First(&problem, problemID); result.Error != nil {
		return nil, result.Error
	}
	return problem.Tags, nil
}

func (s *Storage) SaveTag(ctx context.Context, new models.Tag) (uint, error) {
	if err := s.db.Create(&new).Error; err != nil {
		return 0, err
	}
	return new.ID, nil
}

func (s *Storage) RemoveTag(ctx context.Context, id uint) error {
	if err := s.db.Delete(&models.Tag{}, id).Error; err != nil {
		return err
	}
	return nil
}
