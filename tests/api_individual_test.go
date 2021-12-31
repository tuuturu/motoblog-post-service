package tests

import (
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

			result := bytesToPost(t, recorder.Body.Bytes())

			assert.NotEmpty(t, result.Id)
			assert.NotEmpty(t, result.Time)

			result.Id = ""
			result.Time = ""

			assert.Equal(t, tc.expectResult, &result)
		})
	}
}

func TestUpdatePost(t *testing.T) {
	testCases := []struct {
		name             string
		withExistingPost models.Post
		withUpdate       models.Post

		expectStatus int
		expectResult models.Post
	}{
		{
			name: "Should return 200 and the updated object on successful update",

			withExistingPost: models.Post{
				Title:   "This hear is a post",
				Content: "Not happy with the title",
			},
			withUpdate: models.Post{
				Title:   "I mean, HERE is the post",
				Content: "Thats btetter",
			},
			expectStatus: http.StatusOK,
			expectResult: models.Post{
				Title:   "I mean, HERE is the post",
				Content: "Thats btetter",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			router := core.NewRouter(core.Config{})

			id := createPost(t, router, &tc.withExistingPost)

			request := httptest.NewRequest(
				http.MethodPatch,
				fmt.Sprintf("/posts/%s", id),
				postAsReader(t, tc.withUpdate),
			)
			request.Header.Add("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			assert.Equal(t, tc.expectStatus, recorder.Code)

			result := bytesToPost(t, recorder.Body.Bytes())

			assert.NotEmpty(t, result.Id)
			assert.NotEmpty(t, result.Time)

			result.Id = ""
			result.Time = ""

			assert.Equal(t, tc.expectResult, result)
		})
	}
}

func TestDeletePost(t *testing.T) {
	testCases := []struct {
		name string

		withExistingPost *models.Post
		withIDToDelete   func(string) string

		expectStatus int
	}{
		{
			name: "Should return 204 on successfull deletion",
			withExistingPost: &models.Post{
				Title:   "Not happpy with this post",
				Content: "It barely has any content",
			},
			withIDToDelete: func(original string) string {
				return original
			},
			expectStatus: http.StatusNoContent,
		},
		{
			name:             "Should return 404 on missing post",
			withExistingPost: nil,
			withIDToDelete: func(_ string) string {
				return "eae818ac-a4d8-4a96-bd9c-5ab1dc33039c"
			},
			expectStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var id string
			router := core.NewRouter(core.Config{})

			if tc.withExistingPost != nil {
				id = createPost(t, router, tc.withExistingPost)
			}

			request := httptest.NewRequest(
				http.MethodDelete,
				fmt.Sprintf("/posts/%s", tc.withIDToDelete(id)),
				nil,
			)
			request.Header.Add("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			assert.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}
