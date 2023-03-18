package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arpushkarev/calendar-service/calendar/internal/logger"
)

func StartHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	}

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(config.GetConfig().GetHTTPAddres(), handler) {
		logger.Error.Fatalf("HTTP running failure: %s", err.Error())
	}

	return nil
}
