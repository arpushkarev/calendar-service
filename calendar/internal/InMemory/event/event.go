package event

import (
	"context"

	"github.com/arpushkarev/calendar-service/calendar/internal/model"
)

type InMemory interface {
	Create(ctx context.Context, eventInfo *model.EventInfo) (int64, error)
}

type Service struct {
	info []*model.EventInfo
}

func (s *Service) Create(_ context.Context, eventInfo *model.EventInfo) (int64, error) {
	s.info = append(s.info, eventInfo)
	return int64(len(s.info)) + 1, nil
}
