package entrypoints

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tuuturu/motoblog-post-service/pkg/log"
	"github.com/tuuturu/motoblog-post-service/pkg/models"
	"github.com/tuuturu/motoblog-post-service/pkg/stores"
)

// CreatePost -
func CreatePost(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.GetLogger("entrypoints", "createPost")
		var post models.Post

		err := c.Bind(&post)
		if err != nil {
			logger.Errorf("binding post: %s", err.Error())

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
			logger.Errorf("storing content: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		post.Content = truncateContent(post.Content)

		err = postStore.AddPost(post)
		if err != nil {
			logger.Errorf("storing post: %s", err.Error())

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
