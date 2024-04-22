package models

import (
	"context"
	"gorm.io/gorm"
)

const (
	ProblemDifficultyEasy   = "easy"
	ProblemDifficultyMedium = "medium"
	ProblemDifficultyHard   = "hard"
)

type Problem struct {
	gorm.Model
	Title       string `storage:"not null;unique_index"`
	Description []byte `storage:"type:blob;not null"`
	Difficulty  string `storage:"type:varchar(100);not null"`
	Tags        []Tag  `storage:"many2many:problem_tags;"`
}

type ProblemProvider interface {
	Problem(ctx context.Context, id uint) (*Problem, error)
	ProblemByTitle(ctx context.Context, title string) (*Problem, error)
	ProblemsByFilters(ctx context.Context, difficulty *string, tags []Tag) ([]Problem, error)
}

type ProblemSaver interface {
	Save(
		ctx context.Context,
		title string,
		description []byte,
		difficulty string,
		tags []Tag,
	) (uint, error)
}

type ProblemUpdater interface {
	Update(ctx context.Context, id uint, new Problem) error
}

type ProblemRemover interface {
	Remove(ctx context.Context, id uint) error
}
