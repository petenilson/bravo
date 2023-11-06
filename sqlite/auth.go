package sqlite

import "github.com/petenilson/bravo"

var _ bravo.AuthService = (*AuthService)(nil)

type AuthService struct {
	db *DB
}

func NewAuthService(db *DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

// Create implements bravo.AuthService.
func (*AuthService) Create(*bravo.Auth) error {
	panic("unimplemented")
}

// GetByAthleteId implements bravo.AuthService.
func (*AuthService) GetByAthleteId(athlete_id int) (*bravo.Auth, error) {
	panic("unimplemented")
}

// Update implements bravo.AuthService.
func (*AuthService) Update(auth_id int, update *bravo.AuthUpdate) error {
	panic("unimplemented")
}
