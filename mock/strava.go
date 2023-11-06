package mock

import (
	"time"

	"github.com/petenilson/bravo"
	"github.com/petenilson/bravo/strava"
)

var _ strava.Client = (*StravaClient)(nil)

type StravaClient struct {
	GetEffortFunc   func() (*strava.Effort, error)
}

// GetEfforts implements strava.Client.
func (*StravaClient) GetEfforts(segment_id int, start time.Time, end time.Time) ([]*strava.Effort, error) {
	panic("unimplemented")
}

// NewClient implements strava.Client.
func (*StravaClient) NewClient(*bravo.Auth) (strava.Client, error) {
	panic("unimplemented")
}
