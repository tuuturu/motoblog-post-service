package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deifyed/post-service/pkg/core"
	"github.com/deifyed/post-service/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetIndividualPost(t *testing.T) {
	testCases := []struct {
		name string

		withPost   *models.Post
		withPostID func(string) string

		expectStatus int
		expectResult *models.Post
	}{
		{
			name: "Should retrieve post it created",

			withPost: &models.Post{
				Title:   "Fancy title",
				Content: "what who where which when wazzaaa",
			},
			withPostID: func(original string) string {
				return original
			},
			expectStatus: http.StatusOK,
			expectResult: &models.Post{
				Title:   "Fancy title",
				Content: "what who where which when wazzaaa",
			},
		},
		{
			name: "Should return 404 upon missing post",

			withPost: nil,
			withPostID: func(_ string) string {
				return "4fdee234-8461-4609-b2c2-1d633dc4a6e6"
			},
			expectStatus: http.StatusNotFound,
			expectResult: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			router := core.NewRouter(core.Config{})

			id := createPost(t, router, tc.withPost)

			request := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/posts/%s", tc.withPostID(id)),
				nil,
			)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			assert.Equal(t, tc.expectStatus, recorder.Code)

			if tc.expectResult == nil {
				assert.Empty(t, recorder.Body.Bytes())

				return
			}

			var result models.Post

			err := json.Unmarshal(recorder.Body.Bytes(), &result)
			assert.NoError(t, err)

			assert.NotEmpty(t, result.Id)
			assert.NotEmpty(t, result.Time)

			result.Id = ""
			result.Time = ""

			assert.Equal(t, tc.expectResult, &result)
		})
	}
}
