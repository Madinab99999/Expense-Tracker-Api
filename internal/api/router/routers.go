package router

import (
	"context"
	"net/http"
)

func (r *Router) auth(ctx context.Context) {
	r.mux.HandleFunc("POST /register", r.handler.AuthHandler.Register)
	r.mux.HandleFunc("POST /login", r.handler.AuthHandler.Login)
	r.mux.HandleFunc("POST /access-token", r.handler.AuthHandler.AccessToken)
}

func (r *Router) expense(ctx context.Context) {
	r.mux.Handle("POST /expenses", r.midd.Authenticator(http.HandlerFunc(r.handler.ExpenseHandler.InsertExpense)))
	r.mux.Handle("PUT /expenses/{id}", r.midd.Authenticator(http.HandlerFunc(r.handler.ExpenseHandler.UpdateExpense)))
	r.mux.Handle("DELETE /expenses/{id}", r.midd.Authenticator(http.HandlerFunc(r.handler.ExpenseHandler.DeleteExpense)))
	r.mux.Handle("GET /expenses/{id}", r.midd.Authenticator(http.HandlerFunc(r.handler.ExpenseHandler.GetInformationOfExpense)))
	r.mux.Handle("GET /expenses", r.midd.Authenticator(http.HandlerFunc(r.handler.ExpenseHandler.GetAllExpenses)))
	r.mux.Handle("GET /expenses/stats", r.midd.Authenticator(http.HandlerFunc(r.handler.ExpenseHandler.GetExpenseStats)))
}

func (r *Router) categories(ctx context.Context) {
	r.mux.Handle("GET /categories", r.midd.Authenticator(http.HandlerFunc(r.handler.CategoryHandler.GetAllCategories)))
}
