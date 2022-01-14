package memory

import (
	"github.com/tuuturu/motoblog-post-service/pkg/models"
	"github.com/tuuturu/motoblog-post-service/pkg/stores"
)

func (receiver *store) AddPost(post models.Post) error {
	receiver.db[post.Id] = post

	return nil
}

func (receiver *store) DeletePost(id string) error {
	_, ok := receiver.db[id]
	if !ok {
		return stores.ErrNotFound
	}

	delete(receiver.db, id)

	return nil
}

func (receiver *store) UpdatePost(id string, updatedPost models.Post) error {
	_, ok := receiver.db[id]
	if !ok {
		return stores.ErrNotFound
	}

	receiver.db[id] = updatedPost

	return nil
}

func (receiver *store) GetPost(id string) (models.Post, error) {
	post, ok := receiver.db[id]
	if !ok {
		return models.Post{}, stores.ErrNotFound
	}

	return post, nil
}

func (receiver *store) GetAllPosts(_ stores.Filter) ([]models.Post, error) {
	posts := make([]models.Post, len(receiver.db))
	index := 0

	for _, post := range receiver.db {
		posts[index] = post

		index++
	}

	return posts, nil
}

func New() stores.PostStore {
	return &store{
		db: make(map[string]models.Post),
	}
}
