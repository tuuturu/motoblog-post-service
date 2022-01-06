package entrypoints

import (
	"bytes"
	"net/http"
	"time"

	"github.com/deifyed/post-service/pkg/models"
	"github.com/deifyed/post-service/pkg/stores"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePost -
func CreatePost(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post

		err := c.Bind(&post)
		if err != nil {
			c.Status(http.StatusBadRequest)

			return
		}

		// Validate

		post.Id = uuid.New().String()
		post.Time = time.Now().Format(time.RFC3339)

		err = contentStore.StoreContent(
			post.Id,
			bytes.NewBuffer([]byte(post.Content)),
		)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		post.Content = truncateContent(post.Content)

		err = postStore.AddPost(post)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		c.JSON(http.StatusCreated, post)
	}
}

// GetPosts -
func GetPosts(postStore stores.PostStore, _ stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := stores.Filter{}

		posts, err := postStore.GetAllPosts(filter)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		c.JSON(http.StatusOK, posts)
	}
}
