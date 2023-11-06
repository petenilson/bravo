package bravo

import "time"

type Athlete struct {
	Id        int
	StravaId  int
	FirstName string
	LastName  string
	Auth      *Auth
	AuthId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AthleteService interface {
	Create(*Athlete) error
	Update(*AthleteUpdate) error
	Search(*AthleteFilter) ([]*Athlete, error)
}

type AthleteUpdate struct {
}

type AthleteFilter struct {
}
