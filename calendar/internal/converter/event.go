package converter

import (
	"time"

	"github.com/arpushkarev/calendar-service/calendar/internal/model"
)

func ToModelEventInfo(info *desc.EventInfo) *model.EventInfo {
	return &model.EventInfo{
		Title:        "",
		Date:         time.Time{},
		Duration:     time.Time{},
		Description:  "",
		AuthorID:     0,
		ReminderTime: time.Time{},
	}
}
