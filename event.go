package bravo

import (
	"context"
	"time"
)

type Event struct {
	Id          int
	Name        string
	Segment     *Segment
	SegmentId   int
	Owner       *Athlete
	CreatorId   int
	InviteCode  string
	IsActive    bool
	StartTime   time.Time
	EndTime     time.Time
	Memberships []*Membership
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EventService interface {
	Search(*EventFilter) ([]*Event, error)
}

type EventFilter struct {
	Id     *int
	Active bool
}

type EventMonitor interface {
	Refresh(context.Context) error
}

type EventDetail struct {
	Athletes []*Athlete
}
