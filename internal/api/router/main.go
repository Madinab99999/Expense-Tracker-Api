package router

import (
	"context"
	"net/http"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/handler"
	"github.com/Madinab99999/Expense-Tracker-Api/internal/api/middleware"
)

type Router struct {
	mux     *http.ServeMux
	handler *handler.Handler
	midd    *middleware.Middleware
}

func New(handler *handler.Handler, midd *middleware.Middleware) *Router {
	mux := http.NewServeMux()
	return &Router{mux: mux, handler: handler, midd: midd}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.auth(ctx)
	r.expense(ctx)
	r.categories(ctx)
	return r.mux
}
