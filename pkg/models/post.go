package models

// Post - Schema representing a single post
type Post struct {

	// An ID uniquely identifying a post
	Id string `json:"id,omitempty"`

	// The title of the post
	Title string `json:"title,omitempty"`

	// The date and time the post was submitted in ISO 8601 format
	Time string `json:"time,omitempty"`

	// The content of the post
	Content string `json:"content,omitempty"`

	// The images related to the post
	Images []string `json:"images,omitempty"`
}
