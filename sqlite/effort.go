package sqlite

import "github.com/petenilson/bravo"

var _ bravo.EffortService = (*EffortService)(nil)

type EffortService struct {
	db *DB
}

func NewEffortService(db *DB) *EffortService {
	return &EffortService{db}
}

// Create implements bravo.EffortService.
func (*EffortService) Create(*bravo.Effort) error {
	panic("unimplemented")
}

// GetByID implements bravo.EffortService.
func (*EffortService) GetByID(string) (*bravo.Effort, error) {
	panic("unimplemented")
}
