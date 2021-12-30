package entrypoints

import (
	"github.com/deifyed/post-service/pkg/models"
)

const maxTruncatedLength = 160

func truncateContent(original string) string {
	if len(original) > maxTruncatedLength {
		return original[0:maxTruncatedLength]
	}

	return original
}

func merge(a string, b string) string {
	if a == b {
		return a
	}

	if len(b) > 0 {
		return b
	}

	return a
}

func updatePost(target *models.Post, source models.Post) error {
	target.Title = merge(target.Title, source.Title)
	target.Time = merge(target.Time, source.Time)
	target.Content = merge(target.Content, source.Content)

	return nil
}
