package bravo

import "time"

type Auth struct {
	Id           int
	CreatedAt    time.Time
	ModifiedAt   time.Time
	ExpiresAt    time.Time
	Athlete      *Athlete
	AthleteId    string
	AccessToken  string
	RefreshToken string
}

type AuthService interface {
	Create(*Auth) error
	GetByAthleteId(athlete_id int) (*Auth, error)
	Update(auth_id int, update *AuthUpdate) error
}

type AuthUpdate struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}
