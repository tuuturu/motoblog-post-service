package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deifyed/post-service/pkg/core/entrypoints"
	"github.com/deifyed/post-service/pkg/stores"
	contentMemoryStore "github.com/deifyed/post-service/pkg/stores/content/memory"
	postMemoryStore "github.com/deifyed/post-service/pkg/stores/post/memory"
)

// NewRouter returns a new router.
func NewRouter(_ Config) *gin.Engine {
	router := gin.Default()

	postStore := postMemoryStore.New()
	contentStore := contentMemoryStore.New()

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerGenerator(
				postStore,
				contentStore,
			))
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerGenerator(
				postStore,
				contentStore,
			))
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerGenerator(
				postStore,
				contentStore,
			))
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerGenerator(
				postStore,
				contentStore,
			))
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerGenerator(
				postStore,
				contentStore,
			))
		}
	}

	return router
}

// Index is the index handler.
func Index(_ stores.PostStore, _ stores.ContentStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	}
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"CreatePost",
		http.MethodPost,
		"/posts",
		entrypoints.CreatePost,
	},

	{
		"GetPosts",
		http.MethodGet,
		"/posts",
		entrypoints.GetPosts,
	},

	{
		"DeletePost",
		http.MethodDelete,
		"/posts/:id",
		entrypoints.DeletePost,
	},

	{
		"GetPost",
		http.MethodGet,
		"/posts/:id",
		entrypoints.GetPost,
	},

	{
		"UpdatePost",
		http.MethodPatch,
		"/posts/:id",
		entrypoints.UpdatePost,
	},
}
