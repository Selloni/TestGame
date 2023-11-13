package handler

import (
	"WB/interal"
	"WB/interal/customer/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

// todo: midlle, config

type handler struct {
	ctx  context.Context
	sql  *pgx.Conn
	user interal.Model
}

func NewHandler(ctx context.Context, sql *pgx.Conn, user interal.Model) *handler {
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

func (h *handler) meHandle(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) startHandle(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) loginHandle(w http.ResponseWriter, r *http.Request) {

	var model interal.Model
	model.Login = r.FormValue("login")
	model.Role = r.FormValue("role")
	fmt.Println(model.Role)
	token := interal.GenerateToken(model.Login, model.Role)
	w.Header().Set("Authorization", "Bearer "+token)
	// ищем пользователя в бд и отправляем токен
	http.Error(w, "пользователь или пароль не верен", http.StatusUnauthorized)
}

func (h *handler) registerHandle(w http.ResponseWriter, r *http.Request) {
	// поиск по логину, если нет добавляем в бд
	var model interal.Model
	//todo: валидировать роли
	h.user.Login = r.FormValue("login")
	h.user.Role = r.FormValue("role")
	h.user.Password = r.FormValue("password")
	fmt.Println(model.Role)

	check, err := db.Check(h.ctx, h.sql, h.user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if check {
		http.Error(w, "пользователь уже существует", http.StatusConflict)
		return
	}

	//db.CreateUser(h.ctx, h.sql, h.user)

	token := interal.GenerateToken(model.Login, model.Role)
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) tasksHandler(w http.ResponseWriter, r *http.Request) {
	// бд для задач куда будем сохранять
	task := GenerateTask()
	data, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// для верного интерпретирвания данных
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GenerateTask() (result map[string]int) {
	// либо на месте случайно гененрировать товары
	// либо доставать случайное из бд уже сгенерированные
	return
}
