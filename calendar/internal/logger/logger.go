package logger

import (
	"os"

	"github.com/arpushkarev/calendar-service/calendar/internal/config"
	"github.com/rs/zerolog"
)

type Logger struct {
	Logger *zerolog.Logger
	Level  string
}

func NewLogger(config *config.LoggerConfig) (*Logger, error) {
	file, err := os.OpenFile("calendar.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &Logger{
		Logger: config.LoggerConfig,
	}, nil
}

func (l *Logger) Info(message string, a ...any) {
	l.Logger.Info().Msg(message)
}

func (l *Logger) Warn(message string) {
	l.Logger.Warn().Msg(message)
}

func (l *Logger) Debug(message string) {
	l.Logger.Debug().Msg(message)
}

func (l *Logger) Error(message string) {
	l.Logger.Error().Msg(message)
}
