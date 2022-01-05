package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CorsOptions struct {
	LegalHosts []string
}

func Cors(opts CorsOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodOptions {
			c.Next()

			return
		}

		if !contains(opts.LegalHosts, c.Request.Host) {
			return
		}

		c.Header("Access-Control-Allow-Origin", c.Request.Host)
	}
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
