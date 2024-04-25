package models

import (
	"context"
	"gorm.io/gorm"
)

const (
	ProblemDifficultyEasy   = "Easy"
	ProblemDifficultyMedium = "Medium"
	ProblemDifficultyHard   = "Hard"
)

type Problem struct {
	gorm.Model
	Title       string `gorm:"not null;unique_index"`
	Description []byte `gorm:"type:blob;not null"`
	Difficulty  string `gorm:"type:varchar(100);not null"`
	Tags        []Tag  `gorm:"many2many:problem_tags;association_foreignkey:ID;foreignkey:ID"`
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
