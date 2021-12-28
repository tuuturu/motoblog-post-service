package entrypoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeletePost -
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetPost -
func GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePost -
func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
