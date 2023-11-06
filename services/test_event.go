package services

import (
	"testing"
	"time"

	"github.com/petenilson/bravo"
	"github.com/petenilson/bravo/mock"
	"github.com/petenilson/bravo/strava"
)

func NewTestEventMonitor(
	service bravo.EventService,
	client strava.Client,
) *EventMonitor {
	return &EventMonitor{
		EventService: &service,
		StravaClient: &client,
	}
}

func TestUserHandlers(t *testing.T) {
	athlete := bravo.Athlete{
		FirstName: "Peter",
		LastName:  "Nilson",
		Auth:      &bravo.Auth{},
	}
	membership := bravo.Membership{
		Athlete: &athlete,
	}
	eventService := mock.EventService{
		SearchFunc: func(ef *bravo.EventFilter) ([]*bravo.Event, error) {
			e := bravo.Event{
				ID:          "",
				Name:        "",
				Segment:     &bravo.Segment{},
				SegmentId:   "",
				Owner:       &bravo.Athlete{},
				OwnerId:     "",
				AcCreatorId    false,
				StartTime:   time.Time{},
				EndTime:     time.Time{},
				Memberships: []*bravo.Membership{&membership},
			}
			return []*bravo.Event{&e}, nil
		},
	}
	stravaClient := mock.StravaClient{
		GetEffortFunc: func() (*strava.Effort, error) {
			e := strava.Effort{
				ID:        "",
				SIdrtedAt: time.Time{},
			}
			return &e, nil
		},
	}
	m := NewTestEventMonitor(&eventService, &stravaClient)
	m.Refresh()

}
