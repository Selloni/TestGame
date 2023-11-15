package handler

import (
	"WB/interal"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

type handler struct {
	ctx  context.Context
	sql  *pgxpool.Pool
	user interal.Model
}

func NewHandler(ctx context.Context, sql *pgxpool.Pool, user interal.Model) *handler {
	return &handler{
		ctx:  ctx,
		sql:  sql,
		user: user,
	}
}

func (h *handler) Route() error {

	http.HandleFunc("/login", h.loginHandle)
	http.HandleFunc("/register", h.registerHandle)
	http.HandleFunc("/tasks", h.tasksHandler)
	http.HandleFunc("/me", h.meHandle)
	http.HandleFunc("/start", h.startHandle)

	log.Println("listen port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
