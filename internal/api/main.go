package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/handler"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/middleware"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/repository"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/router"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/service"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/configs"
)

type Api struct {
	logger *slog.Logger
	config *configs.Config
	router *router.Router
	server *http.Server
}

func New(logger *slog.Logger, config *configs.Config, db *sql.DB) *Api {
	repo := repository.New(slog.With("repository", "repository"), db)
	svc := service.New(slog.With("service", "service"), config, repo)
	handler := handler.New(slog.With("handler", "handler"), svc)

	midd := middleware.New(config, slog.With("middleware", "middleware"))

	router := router.New(handler, midd)
	return &Api{
		logger: logger,
		config: config,
		router: router,
	}
}

func (api *Api) Start(ctx context.Context) error {
	mux := api.router.Start(ctx)

	port, err := strconv.Atoi(api.config.ApiPort)
	if err != nil {
		return err
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %d", port)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	api.server = srv

	slog.InfoContext(
		ctx,
		"starting service",
		"port", port,
	)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "service error", "error", err)
		return err
	}

	return nil
}

func (api *Api) Stop(ctx context.Context) error {
	if err := api.server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "server shutdown error", "error", err)
		return err
	}

	return nil
}
