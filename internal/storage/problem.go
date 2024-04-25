package storage

import (
	"context"
	"errors"
	"github.com/ilborsch/leetGo-web/internal/models"
	"github.com/ilborsch/leetGo-web/internal/utils"
	"gorm.io/gorm"
)

func (s *Storage) Problem(ctx context.Context, id uint) (models.Problem, error) {
	var problem models.Problem
	result := s.db.First(&problem, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Problem{}, nil
		}
		return models.Problem{}, result.Error
	}
	return problem, nil
}

func (s *Storage) ProblemByTitle(ctx context.Context, title string) (models.Problem, error) {
	var problem models.Problem
	result := s.db.Where("title = ?", title).First(&problem)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Problem{}, nil
		}
		return models.Problem{}, result.Error
	}
	return problem, nil
}

func (s *Storage) ProblemsByFilters(ctx context.Context, difficulty *string, tags []models.Tag) ([]models.Problem, error) {
	var problems []models.Problem
	if difficulty == nil && len(tags) == 0 {
		return nil, errors.New("invalid filters: both difficulty and tags are empty")
	}

	query := s.db.Model(&models.Problem{})
	if difficulty != nil {
		query = query.Where("difficulty = ?", *difficulty)
	}
	if len(tags) > 0 {
		tagIDs := utils.GetTagIDs(tags)
		query = query.Joins("JOIN problem_tags ON problem.id = problem_tags.problem_id").
			Where("problem_tags.tag_id IN ?", tagIDs)
	}

	if err := query.Find(&problems).Error; err != nil {
		return nil, err
	}
	return problems, nil
}

func (s *Storage) Problems(ctx context.Context) ([]models.Problem, error) {
	var problems []models.Problem
	if result := s.db.Find(&problems); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return problems, nil
}

func (s *Storage) SaveProblem(ctx context.Context, new models.Problem) (uint, error) {
	if err := s.db.Create(&new).Error; err != nil {
		return 0, err
	}
	return new.ID, nil
}

func (s *Storage) RemoveProblem(ctx context.Context, id uint) error {
	if err := s.db.Delete(&models.Article{}, id).Error; err != nil {
		return err
	}
	return nil
}
