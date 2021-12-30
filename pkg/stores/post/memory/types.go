package memory

import "github.com/deifyed/post-service/pkg/models"

type store struct {
	db map[string]models.Post
}
