package services

import (
	"context"

	"github.com/petenilson/bravo"
	"github.com/petenilson/bravo/http"
	"github.com/petenilson/bravo/sqlite"
)

var _ bravo.EventMonitor = (*EventMonitor)(nil)

type EventMonitor struct {
	EventService  bravo.EventService
	AuthService   bravo.AuthService
	EffortService bravo.EffortService
}

func NewEventMonitor(db *sqlite.DB) *EventMonitor {
	return &EventMonitor{
		EventService:  sqlite.NewEventService(db),
		AuthService:   sqlite.NewAuthService(db),
		EffortService: sqlite.NewEffortService(db),
	}
}

// Refresh implements bravo.EventMonitor.
func (m *EventMonitor) Refresh(ctx context.Context) error {
	events, err := m.EventService.Search(&bravo.EventFilter{Active: true})
	if err != nil {
		return err
	}
	a := make(map[*bravo.Athlete][]*bravo.Event)
	for _, event := range events {
		for _, membership := range event.Memberships {
			a[membership.Athlete] = append(a[membership.Athlete], event)
		}
	}
	// for every athlete
	for k, v := range a {
		client := http.NewClient(k, m.AuthService)
		if err != nil {
			return err
		}
		efforts := []*bravo.Effort{}
		// for every event the athlete belongs to
		for _, event := range v {
			e, err := client.GetEfforts(ctx, event.SegmentId, event.StartTime, event.EndTime)
			if err != nil {
				return err
			}
			// for every effort pulled from strava
			for _, v := range e {
				efforts = append(efforts, &bravo.Effort{
					StravaId:  v.Id,
					EventId:   event.Id,
					SegmentId: event.SegmentId,
					StartDate: v.StartDateLocal,
				})
			}
		}
		for _, effort := range efforts {
			m.EffortService.Create(effort)
		}
	}

	return nil
}

func (m *EventMonitor) Start() error {
	panic("unimplemented")
}

func (m *EventMonitor) Stop() error {
	panic("unimplemented")
}
