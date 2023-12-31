// This file was auto-generated by Fern from our API Definition.

package geo

// Geographical coordinates for a location on Planet Earth.
type Coord struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Distance available in multiple units.
type Distance struct {
	Kilometers float64 `json:"kilometers"`
	Miles      float64 `json:"miles"`
}
