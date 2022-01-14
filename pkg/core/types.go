package core

import (
	"github.com/gin-gonic/gin"
	"github.com/tuuturu/motoblog-post-service/pkg/stores"
)

type getFn func(string) string

type Config struct {
	Port       int
	LegalHosts []string
}

type HandlerGeneratorFn func(stores.PostStore, stores.ContentStore) gin.HandlerFunc

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerGenerator is the handler function of this route.
	HandlerGenerator HandlerGeneratorFn
}

// Routes is the list of the generated Route.
type Routes []Route
