package handler

import (
	"WB/interal"
	"WB/interal/customer"
	dbCustomer "WB/interal/customer/db"
	"WB/interal/loader"
	dbLoader "WB/interal/loader/db"
	"WB/interal/posgresql"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

// todo: midlle, config

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

func (h *handler) meHandle(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) startHandle(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) loginHandle(w http.ResponseWriter, r *http.Request) {

	h.user.Login = r.FormValue("login")
	h.user.Role = r.FormValue("role")
	fmt.Println(h.user.Role)
	token := interal.GenerateToken(h.user.Login, h.user.Role)
	w.Header().Set("Authorization", "Bearer "+token)
	// ищем пользователя в бд и отправляем токен
	http.Error(w, "пользователь или пароль не верен", http.StatusUnauthorized)
}

func (h *handler) registerHandle(w http.ResponseWriter, r *http.Request) {
	// поиск по логину, если нет добавляем в бд

	h.user.Login = r.FormValue("login")
	h.user.Role = r.FormValue("role")
	h.user.Password = r.FormValue("password")

	check, err := posgresql.Check(h.ctx, h.sql, h.user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if check {
		http.Error(w, "пользователь уже существует", http.StatusConflict)
		return
	}
	if h.user.Role == "customer" {
		fmt.Println(customer.GenerateCustomer())
		h.user.Customer = customer.GenerateCustomer()
		dbCustomer.CreateUser(h.ctx, h.sql, h.user)
	} else if h.user.Role == "loader" {
		fmt.Println(loader.GenerateLoader())
		h.user.Loader = loader.GenerateLoader()
		dbLoader.CreateUser(h.ctx, h.sql, h.user)
	} else {
		http.Error(w, "достпуные роли  customer & loader", http.StatusBadRequest)
		return
	}

	token := interal.GenerateToken(h.user.Login, h.user.Role)
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusCreated)

}

func (h *handler) tasksHandler(w http.ResponseWriter, r *http.Request) {
	// бд для задач куда будем сохранять
	task := h.GenerateTask()
	data, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// для верного интерпретирвания данных
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *handler) GenerateTask() map[string]int {
	// либо на месте случайно гененрировать товары
	// либо доставать случайное из бд уже сгенерированные
	var mm = map[string]int{
		"bead":   18,
		"fridge": 21,
		"TV":     12,
		"spoon":  2,
	}
	return mm
}
