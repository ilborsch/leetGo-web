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

// ProblemRaw used to initialize problem object which will be later used for create/update operations
func ProblemRaw(title string, description []byte, difficulty string, tags []Tag) Problem {
	return Problem{
		Title:       title,
		Description: description,
		Difficulty:  difficulty,
		Tags:        tags,
	}
}

type ProblemProvider interface {
	Problem(ctx context.Context, id uint) (Problem, error)
	ProblemByTitle(ctx context.Context, title string) (Problem, error)
	ProblemsByFilters(ctx context.Context, difficulty *string, tags []Tag) ([]Problem, error)
	Problems(ctx context.Context) ([]Problem, error)
}

type ProblemSaver interface {
	SaveProblem(ctx context.Context, new Problem) (uint, error)
}

type ProblemRemover interface {
	RemoveProblem(ctx context.Context, id uint) error
}
