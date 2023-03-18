package model

import "time"

type EventInfo struct {
	Title        string    `db:"title"`
	Date         time.Time `db:"date"`
	Duration     time.Time `db:"duration"`
	Description  string    `db:"description"`
	AuthorID     int64     `db:"author_ID"`
	ReminderTime time.Time `db:"reminder_time"`
}

type Event struct {
	Id    int64
	Event *EventInfo
}
