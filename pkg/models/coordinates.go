package models

// Coordinates - Schema representing a geographical point
type Coordinates struct {

	// Latitude defines the latitude of the point
	Latitude float32 `json:"latitude,omitempty"`

	// Longitude defines the Longitude of the point
	Longitude float32 `json:"longitude,omitempty"`
}
