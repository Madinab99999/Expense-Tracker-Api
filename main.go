package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/api"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/db"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/db/repository"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/service"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/configs"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := configs.LoadConfig()

	db, err := db.New(slog.With("db", "db"), cfg)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"initialize service error",
			"service", "db",
			"error", err,
		)
		panic(err)
	}

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	if err := db.Init(ctx); err != nil {
		panic(err)
	}
	repo := repository.New(slog.With("repository", "repository"), db.Pg)
	svc := service.New(slog.With("service", "service"), cfg, repo)
	a := api.New(slog.With("api", "api"), cfg, db.Pg, svc)
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := a.Start(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to start api", "error", err.Error())
		}

		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancelFunc()
	}(cancel)

	<-ctx.Done()

	if err := a.Stop(ctx); err != nil {
		slog.ErrorContext(ctx, "service stop error", "error", err)
	}

	slog.InfoContext(ctx, "server was successfully shutdown.")
}
