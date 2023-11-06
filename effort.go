package bravo

import "time"

type Effort struct {
	Id        int
	StravaId  int
	Event     *Event
	EventId   int
	Segment   *Segment
	SegmentId int
	Athlete   *Athlete
	AthleteId int
	StartDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EffortService interface {
	Create(*Effort) error
	GetByID(string) (*Effort, error)
}

func (e *Effort) Validate() bool {
	if e.StartDate.After(e.Event.EndTime) || e.StartDate.Before(e.Event.StartTime) {
		return false
	}
	return true
}
