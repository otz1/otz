package entity

import "time"

// MeasurementDetail is information that provides
// how long the request took and how many results were given
type MeasurementDetail struct {
	// perhaps elapsed time could be a float?
	ElapsedTime time.Duration `json:"elapsedTime"`
	ResultCount int           `json:"resultCount"`
}
