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
	//fmt.Println(r.Header.Get("Authorization"))
	var x any = h.ctx.Value("token")
	if x == nil {
		http.Error(w, "не найден токен авторизации", http.StatusUnauthorized)
		return
	}
	isValid, _, role := interal.ValidateToken(x.(string))
	if !isValid {
		fmt.Println(isValid)
		http.Error(w, "не найден токен авторизации", http.StatusUnauthorized)
		return
	}
	fmt.Println(role)
	if role == "customer" {
		fmt.Println(h.user.Customer)
	} else if role == "loader" {
		fmt.Println(h.user)
	}
}

func (h *handler) startHandle(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) loginHandle(w http.ResponseWriter, r *http.Request) {
	err, exists := h.existsUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "пользователь не найден или неверный пароль", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Привет, ты сегодня чудестно выглядишь "))
	token, err := interal.GenerateToken(h.user.Login, h.user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.ctx = context.WithValue(h.ctx, "token", token)

	w.WriteHeader(http.StatusOK)

	//fmt.Fprintf(w, token)
	//w.Write([]byte(token))
	//w.Header().Set("Authorization", "Bearer "+token)
	// w - каждый запрос новый w, и так пртсо передать через него не получиться
	//fmt.Println(w.Header().Get("Authorization"))
	// ищем пользователя в бд и отправляем токен
}

func (h *handler) registerHandle(w http.ResponseWriter, r *http.Request) {
	h.user.Login = r.FormValue("login")
	h.user.Role = r.FormValue("role")
	h.user.Password = r.FormValue("password")

	if h.user.Role == "customer" {
		h.user.Customer = customer.GenerateCustomer()
		if err := dbCustomer.CreateUser(h.ctx, h.sql, h.user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "customer(заказчик) успешно создан")

	} else if h.user.Role == "loader" {
		h.user.Loader = loader.GenerateLoader()
		if err := dbLoader.CreateUser(h.ctx, h.sql, h.user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "loader(грузчик) успешно создан")

	} else {
		http.Error(w, "достпуные роли  customer & loader", http.StatusBadRequest)
		return
	}
	//token, err := interal.GenerateToken(h.user.Login, h.user.Role)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusCreated)

}

func (h *handler) tasksHandler(w http.ResponseWriter, r *http.Request) {
	// бд для задач куда будем сохранять
	task := h.generateTask()
	data, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// для верного интерпретирвания данных
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *handler) generateTask() map[string]int {
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

func (h *handler) existsUser(r *http.Request) (error, bool) {
	//todo: проверить на заполнение трех полей
	h.user.Login = r.FormValue("login")
	h.user.Role = r.FormValue("role")
	h.user.Password = r.FormValue("password")

	check, err := posgresql.Check(h.ctx, h.sql, h.user)

	if err != nil {
		return err, false
	}

	return nil, check
}
