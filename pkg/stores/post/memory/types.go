package memory

import "github.com/tuuturu/motoblog-post-service/pkg/models"

type store struct {
	db map[string]models.Post
}
