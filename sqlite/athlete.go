package sqlite

import "github.com/petenilson/bravo"

var _ bravo.AthleteService = (*AthleteService)(nil)

type AthleteService struct {
	db *DB
}

// Search implements bravo.AthleteService.
func (*AthleteService) Search(*bravo.AthleteFilter) ([]*bravo.Athlete, error) {
	panic("unimplemented")
}

// Create implements bravo.AthleteService.
func (*AthleteService) Create(*bravo.Athlete) error {
	panic("unimplemented")
}

// Update implements bravo.AthleteService.
func (*AthleteService) Update(*bravo.AthleteUpdate) error {
	panic("unimplemented")
}
