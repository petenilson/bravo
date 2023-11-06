package strava

import (
	"github.com/petenilson/bravo"
	"time"
)

type Client interface {
	GetEfforts(segment_id int, start, end time.Time) ([]*Effort, error)
	NewClient(*bravo.Auth) (Client, error)
}

type Authenticator interface {
	Authenticate(*bravo.Auth) error
}

