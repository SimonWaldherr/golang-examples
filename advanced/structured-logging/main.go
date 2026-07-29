// Description: Structured logging with log/slog
// Tags: log, slog, structured logging, JSON, context
package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(_ []string, attribute slog.Attr) slog.Attr {
			// Removing the timestamp makes this example's output reproducible.
			if attribute.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return attribute
		},
	})

	logger := slog.New(handler).With(
		slog.String("service", "example-api"),
		slog.String("version", "1.0.0"),
	)

	ctx := context.Background()
	logger.InfoContext(ctx, "request completed",
		slog.String("request_id", "req-42"),
		slog.Duration("duration", 125_000_000),
		slog.Group("http",
			slog.String("method", "GET"),
			slog.Int("status", 200),
		),
	)
}
