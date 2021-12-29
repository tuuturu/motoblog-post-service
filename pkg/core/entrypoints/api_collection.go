package entrypoints

import (
	"net/http"

	"github.com/deifyed/post-service/pkg/stores"
	"github.com/gin-gonic/gin"
)

// CreatePost -
func CreatePost(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}

// GetPosts -
func GetPosts(postStore stores.PostStore, contentStore stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}
