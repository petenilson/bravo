package sqlite

import (
	"context"
	"strings"

	"github.com/petenilson/bravo"
)

var _ bravo.EventService = (*EventService)(nil)

type EventService struct {
	db *DB
}

func NewEventService(db *DB) *EventService {
	return &EventService{
		db: &DB{},
	}
}

// Search implements bravo.EventService.
func (*EventService) Search(*bravo.EventFilter) ([]*bravo.Event, error) {
	panic("unimplemented")
}

func findEvents(ctx context.Context, tx *Tx, filter *bravo.EventFilter) ([]*bravo.Event, int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	// filter by is active
	boolMapping := map[bool]int{true: 1, false: 0}
	where, args = append(where, "active = ?"), append(args, boolMapping[filter.Active])
	// filter by EventId
	if v := filter.Id; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	// Execue query with limiting WHERE clause and LIMIT/OFFSET injected.
	rows, err := tx.QueryContext(ctx, `
		SELECT 
		    id,
		    name,
		    segment_id,
		    creator_id,
		    invite_code,
		    is_active,
		    start_time,
				end_time,
		    created_at,
		    updated_at,
		    COUNT(*) OVER()
		FROM dials
		WHERE `+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`,
		args...,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	events := make([]*bravo.Event, 0)
	num_events := 0
	for rows.Next() {
		var event bravo.Event
		if err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.SegmentId,
			&event.CreatorId,
			&event.InviteCode,
			&event.IsActive,
			(*NullTime)(&event.StartTime),
			(*NullTime)(&event.EndTime),
			(*NullTime)(&event.CreatedAt),
			(*NullTime)(&event.UpdatedAt),
			&num_events,
		); err != nil {
			return nil, 0, err
		}
		events = append(events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return events, num_events, nil
	
}
