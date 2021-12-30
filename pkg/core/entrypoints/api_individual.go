package entrypoints

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/deifyed/post-service/pkg/models"
	"github.com/deifyed/post-service/pkg/stores"
	"github.com/gin-gonic/gin"
)

// DeletePost -
func DeletePost(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := postStore.DeletePost(id)
		if err != nil {
			log.Printf("unable to delete post: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		err = contentStore.DeleteContent(id)
		if err != nil {
			log.Printf("unable to delete content: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		c.Status(http.StatusNoContent)
	}
}

// GetPost -
func GetPost(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		post, err := postStore.GetPost(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		content, err := contentStore.RetrieveContent(id)
		if err != nil {
			status := http.StatusInternalServerError

			if errors.Is(err, stores.ErrNotFound) {
				status = http.StatusNotFound
			}

			c.Status(status)

			return
		}

		raw, err := io.ReadAll(content)
		if err != nil {
			log.Printf("Error buffering content: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		post.Content = string(raw)

		c.JSON(http.StatusOK, post)
	}
}

// UpdatePost -
func UpdatePost(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatedPost models.Post

		err := c.Bind(&updatedPost)
		if err != nil {
			log.Printf("Error binding post: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		originalPost, err := postStore.GetPost(updatedPost.Id)
		if err != nil {
			log.Printf("Error retrieving original post: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		updatePost(&originalPost, updatedPost)

		err = contentStore.StoreContent(
			originalPost.Id,
			bytes.NewReader([]byte(originalPost.Content)),
		)
		if err != nil {
			log.Printf("Error updating content: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		originalPost.Content = truncateContent(originalPost.Content)

		err = postStore.UpdatePost(originalPost.Id, originalPost)
		if err != nil {
			log.Printf("Error updating post: %s", err.Error())

			c.Status(http.StatusInternalServerError)

			return
		}

		c.JSON(http.StatusOK, originalPost)
	}
}
