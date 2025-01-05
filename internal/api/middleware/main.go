package middleware

import (
	"log/slog"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/configs"
)

type Middleware struct {
	config *configs.Config
	log    *slog.Logger
}

func New(config *configs.Config, log *slog.Logger) *Middleware {
	return &Middleware{config: config, log: log}
}
