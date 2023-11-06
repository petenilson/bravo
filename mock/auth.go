package mock

import (
	"github.com/petenilson/bravo"
)

var _ bravo.AuthService = (*AuthService)(nil)

type AuthService struct {
	CreateFunc         func(*bravo.Auth) error
	GetByAthleteIdFunc func(int) (*bravo.Auth, error)
	UpdateFunc         func(int, *bravo.AuthUpdate) error
}

// Create implements bravo.AuthService.
func (s *AuthService) Create(a *bravo.Auth) error {
	return s.CreateFunc(a)
}

// GetByAthleteId implements bravo.AuthService.
func (s *AuthService) GetByAthleteId(athlete_id int) (*bravo.Auth, error) {
	return s.GetByAthleteIdFunc(athlete_id)
}

// Update implements bravo.AuthService.
func (s *AuthService) Update(auth_id int, update *bravo.AuthUpdate) error {
	return s.UpdateFunc(auth_id, update)
}
