package bravo

import "time"

type Segment struct {
	StravaId int
	Name     string
	City     string
	Country  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
