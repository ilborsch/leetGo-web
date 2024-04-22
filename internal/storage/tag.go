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

func (s *Storage) TagByName(ctx context.Context, name string) (models.Tag, error) {
	var tag models.Tag
	result := s.db.Where("name = ?", name).First(&tag)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Tag{}, nil
		}
		return models.Tag{}, result.Error
	}
	return tag, nil
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
