package bravo

import "time"

type Membership struct {
	DateCreated time.Time
	Event       *Event
	EventId     string
	Athlete     *Athlete
	AthleteId   string
}

type MembershipService interface {
	Search(*MembershipFilter) ([]*Membership, error)
}

type MembershipFilter struct {
	EventId string
}
