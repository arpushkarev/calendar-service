package db

import (
	"context"
	//"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/arpushkarev/calendar-service/calendar/internal/model"
	"github.com/arpushkarev/calendar-service/calendar/internal/pkg/db"
)

const (
	tableName = "calendar"
)

type repository struct {
	client db.Client
}

func NewRepository(client db.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, eventInfo *model.EventInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "date_event", "duration", "description", "author_ID", "reminder_time").
		Values(
			eventInfo.Title,
			eventInfo.Date,
			eventInfo.Duration,
			eventInfo.Description,
			eventInfo.AuthorID,
			eventInfo.ReminderTime,
		).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "Create",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()

	var id int64

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//func (r *repository) Update(ctx context.Context, req *model.UpdateEventInfo) error {
//
//}

//func (r *repository) Delete(ctx context.Context, id int64) error {
//
//}

//func (r *repository) Get(ctx context.Context, date time.Time) (*model.Event, error) {
//
//}

//func (r *repository) GetWeek(ctx context.Context, week time.Time) ([]*model.Event, error) {
//
//}

//func (r *repository) GetMonth(ctx context.Context, month time.Time) ([]*model.Event, error) {
//
//}
