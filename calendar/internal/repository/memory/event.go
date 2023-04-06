package memory

import (
	"context"

	"github.com/arpushkarev/calendar-service/calendar/internal/model"
)

type repository struct {
	memo map[int64]*model.EventInfo
}

func NewRepository() *repository {
	memo := make(map[int64]*model.EventInfo)
	return &repository{
		memo: memo,
	}
}

func (r *repository) Create(ctx context.Context, eventInfo *model.EventInfo) (int64, error) {
	key := int64(len(r.memo))

	key += 1

	r.memo[key] = eventInfo

	return key, nil
}
