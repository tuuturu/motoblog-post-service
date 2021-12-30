package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deifyed/post-service/pkg/core"
	"github.com/deifyed/post-service/pkg/models"
)

func TestCreatePost(t *testing.T) {
	testCases := []struct {
		name string

		withPosts []models.Post

		expectStatus []int
	}{
		{
			name: "Should return 204 upon successfully created post",
			withPosts: []models.Post{
				{
					Title:   "Dummy",
					Content: "Just some dummy text bruh",
				},
			},
			expectStatus: []int{http.StatusCreated},
		},
	}

	for _, tc := range testCases {
		router := core.NewRouter(core.Config{})

		for index, post := range tc.withPosts {
			raw, err := json.Marshal(post)
			assert.NoError(t, err)

			request := httptest.NewRequest(
				http.MethodPost,
				"/posts",
				bytes.NewReader(raw),
			)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			assert.Equal(t, tc.expectStatus[index], recorder.Code)
		}
	}
}
