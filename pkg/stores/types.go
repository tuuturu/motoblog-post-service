package stores

import (
	"errors"
	"github.com/deifyed/post-service/pkg/models"
	"io"
)

var ErrNotFound = errors.New("not found")

type Filter struct{}

type PostStore interface {
	AddPost(models.Post) error
	DeletePost(id string) error
	UpdatePost(id string, updatedPost models.Post) error
	GetPost(id string) (models.Post, error)
	GetAllPosts(Filter) ([]models.Post, error)
}

type ContentStore interface {
	StoreContent(id string, content io.Reader) error
	RetrieveContent(id string) (io.Reader, error)
	DeleteContent(id string) error
}
