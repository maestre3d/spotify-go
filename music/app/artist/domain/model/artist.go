package model

import "time"

// Artist Artist entity model
type Artist struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Birth            time.Time `json:"birth"`
	TotalFollowers   int64     `json:"total_followers"`
	MonthlyListeners int64     `json:"monthly_listeners"`
	Active           bool      `json:"active"`
	Nationality      string    `json:"nationality"`
	ImageURL         *string   `json:"image_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
