package app

import (
	"context"
	"net/http"

	"github.com/arpushkarev/calendar-service/calendar/internal/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type App struct {
	Logger          *logger.Logger
	serviceProvider *serviceProvider
	pathConfig      string
	mux             *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}

	err := a.initDeps(ctx)

	return a, err
}

func (a *App) StartHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := http.ListenAndServe(a.serviceProvider.GetConfig().GetHTTPAddress(), a.mux); err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	if a.Logger == nil {
		lgr, err := logger.NewLogger(a.serviceProvider.GetConfig().GetLoggerConfig())
		if err != nil {
			logger.Logger.Info("Failed to init logger", err.Error())
		}

		a.Logger = lgr
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initLogger,
		//a.initHTTPServer,
		//a.initGRPCServer,
		//a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

//func (a *App) initServer (ctx context.Context) error{
//
//}

//func (a *App) initHTTPServer(ctx context.Context) error {
//
//}

//func (a *App) initGRPCServer(ctx context.Context) error {
//
//}

func (a *App) Run() error {
	err := a.StartHTTP()
	if err != nil {
		return err
	}

	return nil
}
