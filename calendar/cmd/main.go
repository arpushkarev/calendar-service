package main

import (
	"context"
	"flag"

	"github.com/arpushkarev/calendar-service/calendar/internal/app"
	"github.com/arpushkarev/calendar-service/calendar/internal/logger"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "config/config.json", "Path to configuration")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		logger.Logger.Info("New app creation failure:", err.Error())
	}

	err = a.Run()
	if err != nil {
		logger.Error.Fatalf("App running failure: %s", err.Error())
	}

}
