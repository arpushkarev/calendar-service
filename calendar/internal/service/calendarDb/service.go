package calendarDb

import "github.com/arpushkarev/calendar-service/calendar/internal/repository/event"

type Service struct {
	eventRepository event.Repository
}

func NewService(eventRepository event.Repository) *Service {
	return &Service{
		eventRepository: eventRepository,
	}
}
