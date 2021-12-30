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
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			router := core.NewRouter(core.Config{})

			for index, post := range tc.withPosts {
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

				assert.Equal(t, tc.expectStatus[index], recorder.Code)
			}
		})
	}
}

func TestGetAllPosts(t *testing.T) {
	testCases := []struct {
		name string

		withPosts   []models.Post
		expectPosts []models.Post
	}{
		{
			name: "Should have one post",
			withPosts: []models.Post{
				{
					Title:   "Dummy",
					Content: "Dummy content",
				},
			},
			expectPosts: []models.Post{
				{
					Title:   "Dummy",
					Content: "Dummy content",
				},
			},
		},
		{
			name: "Should have multiple posts",
			withPosts: []models.Post{
				{
					Title:   "Dummy",
					Content: "Dummy content",
				},
				{
					Title:   "Dummy number two",
					Content: "more Dummy content",
				},
				{
					Title:   "Dummy number three",
					Content: "even more Dummy content",
				},
			},
			expectPosts: []models.Post{
				{
					Title:   "Dummy",
					Content: "Dummy content",
				},
				{
					Title:   "Dummy number two",
					Content: "more Dummy content",
				},
				{
					Title:   "Dummy number three",
					Content: "even more Dummy content",
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			router := core.NewRouter(core.Config{})

			for _, post := range tc.withPosts {
				createPost(t, router, &post)
			}

			request := httptest.NewRequest(http.MethodGet, "/posts", nil)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			assert.Equal(t, http.StatusOK, recorder.Code)

			var returnedPosts []models.Post

			err := json.Unmarshal(recorder.Body.Bytes(), &returnedPosts)
			assert.NoError(t, err)

			for index := range returnedPosts {
				assert.NotEmpty(t, returnedPosts[index].Id)
				assert.NotEmpty(t, returnedPosts[index].Time)

				returnedPosts[index].Id = ""
				returnedPosts[index].Time = ""
			}

			for _, post := range tc.expectPosts {
				assert.Contains(t, returnedPosts, post)
			}
		})
	}
}
