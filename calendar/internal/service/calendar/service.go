package calendar

import "github.com/arpushkarev/calendar-service/calendar/internal/repository"

type Service struct {
	eventRepository repository.Repository
}

func NewService(eventRepository repository.Repository) *Service {
	return &Service{
		eventRepository: eventRepository,
	}
}
