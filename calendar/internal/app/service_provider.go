package app

import (
	"context"
	"log"

	"github.com/arpushkarev/calendar-service/calendar/internal/config"
	"github.com/arpushkarev/calendar-service/calendar/internal/logger"
	"github.com/arpushkarev/calendar-service/calendar/internal/pkg/db"
	"github.com/arpushkarev/calendar-service/calendar/internal/repository"
	rdb "github.com/arpushkarev/calendar-service/calendar/internal/repository/db"
	"github.com/arpushkarev/calendar-service/calendar/internal/repository/memory"
	"github.com/arpushkarev/calendar-service/calendar/internal/service/calendar"
)

type serviceProvider struct {
	db              db.Client
	configPath      string
	config          config.IConfig
	logger          *logger.Logger
	repository      repository.Repository
	calendarService *calendar.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{configPath: configPath}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDbConfig()
		if err != nil {
			s.logger.Info("Failed to get db config:", err)
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			s.logger.Info("Failed to create db client:", err)
		}

		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() config.IConfig {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			s.logger.Info("Failed to get config:", err)
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetLogger() *logger.Logger {
	if s.logger == nil {
		lgr, err := logger.NewLogger(s.GetConfig().GetLoggerConfig())
		if err != nil {
			log.Fatalf("Failed to get logger: %s", err.Error())
		}
		s.logger = lgr
	}

	return s.logger
}

func (s *serviceProvider) GetRepository(ctx context.Context) repository.Repository {
	if s.repository == nil {
		if s.GetConfig().GetDataSource().Repos == "db" {
			s.repository = rdb.NewRepository(s.GetDB(ctx))
		} else if s.GetConfig().GetDataSource().Repos == "memory" {
			s.repository = memory.NewRepository()
		}
	}

	return s.repository
}

func (s *serviceProvider) GetCalendarService(ctx context.Context) *calendar.Service {
	if s.calendarService == nil {
		s.calendarService = calendar.NewService(s.GetRepository(ctx))
	}

	return s.calendarService
}
