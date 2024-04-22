package utils

import "github.com/ilborsch/leetGo-web/internal/models"

func GetTagNames(tags []models.Tag) []string {
	tagNames := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	return tagNames
}
