package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tuuturu/motoblog-post-service/pkg/models"
)

func createPost(t *testing.T, router *gin.Engine, post *models.Post) string {
	if post == nil {
		return "noID"
	}

	raw, err := json.Marshal(post)
	assert.NoError(t, err)

	request := httptest.NewRequest(
		http.MethodPost,
		"/posts",
		bytes.NewReader(raw),
	)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var result models.Post

	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.NoError(t, err)

	return result.Id
}

func postAsReader(t *testing.T, post models.Post) io.Reader {
	raw, err := json.Marshal(post)
	assert.NoError(t, err)

	return bytes.NewReader(raw)
}

func bytesToPost(t *testing.T, raw []byte) models.Post {
	var result models.Post

	err := json.Unmarshal(raw, &result)
	assert.NoError(t, err)

	return result
}
