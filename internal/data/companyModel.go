package data

import (
	"time"
)

type Company struct {
	ID            int64     `json:"id"`         // Unique integer ID for the company
	CreatedAt     time.Time `json:"created_at"` // Timestamp for when the movie is added to db
	Name          string    `json:"name"`
	Year          int32     `json:"year"` // When was it founded
	Valuation     int64     `json:"valuation"`
	Headcount     int64     `json:"headcount"`      // Number of employees
	IndustryField []string  `json:"industry_field"` // type of industry
	Timestamp     time.Time `json:"timestamp"`      // the time when the data was gathered
	Version       int32     `json:"version"`
}
