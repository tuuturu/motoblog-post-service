package entrypoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost -
func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetPosts -
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
