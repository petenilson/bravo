package mock

import "github.com/petenilson/bravo"

var _ bravo.EventService = (*EventService)(nil)

type EventService struct {
  SearchFunc func(*bravo.EventFilter) ([]*bravo.Event, error)
}

// Search implements bravo.EventService.
func (s *EventService) Search(filter *bravo.EventFilter) ([]*bravo.Event, error) {
  return s.SearchFunc(filter)
}
