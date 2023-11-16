package handler

import (
	"WB/interal"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"time"
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

	http.HandleFunc("/login", middleware(h.loginHandle))
	http.HandleFunc("/register", middleware(h.registerHandle))
	http.HandleFunc("/tasks", middleware(h.tasksHandler))
	http.HandleFunc("/me", middleware(h.meHandle))
	http.HandleFunc("/start", middleware(h.startHandle))

	log.Println("listen port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request: %s, %s\n", time.Now(), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
