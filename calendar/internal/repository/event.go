package repository

import (
	"context"

	"github.com/arpushkarev/calendar-service/calendar/internal/model"
)

type Repository interface {
	Create(ctx context.Context, eventInfo *model.EventInfo) (int64, error)
	//Update(ctx context.Context, req *model.UpdateEventInfo) error
	//Delete(ctx context.Context, id int64) error
	//Get(ctx context.Context, date time.Time) (*model.Event, error)
	//GetWeek(ctx context.Context, week time.Time) ([]*model.Event, error)
	//GetMonth(ctx context.Context, month time.Time) ([]*model.Event, error)
}
