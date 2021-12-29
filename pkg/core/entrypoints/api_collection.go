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
		post.Time = time.Now().String()

		err = postStore.AddPost(post)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		err = contentStore.StoreContent(
			post.Id,
			bytes.NewBuffer([]byte(post.Content)),
		)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		c.JSON(http.StatusCreated, gin.H{})
	}
}

// GetPosts -
func GetPosts(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}
