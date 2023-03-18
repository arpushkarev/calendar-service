package calendarDb

import (
	"context"

	"github.com/arpushkarev/calendar-service/calendar/internal/model"
)

func (s *Service) Create(ctx context.Context, eventInfo *model.EventInfo) (int64, error) {
	id, err := s.eventRepository.Create(ctx, eventInfo)
	if err != nil {
		return 0, err
	}

	return id, nil
}
