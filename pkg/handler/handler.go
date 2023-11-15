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
	var x any = h.ctx.Value("token")
	if x == nil {
		http.Error(w, "не найден токен авторизации", http.StatusUnauthorized)
		return
	}
	isValid, _, role := interal.ValidateToken(x.(string))
	if !isValid {
		http.Error(w, "не валидный токен", http.StatusUnauthorized)
		return
	}
	if role == "customer" {
		str := fmt.Sprintf("Привет %s, твой бюджет %d.\nТы можешь нанять этих ребят\n",
			h.user.Login, h.user.Customer.Money)
		allLoader, err := dbLoader.GetAllLoader(h.ctx, h.sql)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, str, "\n")
		fmt.Fprint(w, "id вес ЗП Пьет Устал\n")
		for i := range allLoader {
			fmt.Fprint(w, allLoader[i], "\n")
		}

	} else if role == "loader" {
		str := fmt.Sprintf("Привет %s твоя анкета:\nМакс вес - %v\nЗП - "+
			"%d\nВредные привычки - %t\nУсталость - %d\n",
			h.user.Login, h.user.Loader.Weight, h.user.Loader.Salary,
			h.user.Loader.Drunk, h.user.Loader.Tired)
		fmt.Fprint(w, str)
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

	if h.user.Role == "customer" {
		user, err := dbCustomer.GetInfo(h.ctx, h.sql, h.user.Login)
		if err != nil {
			log.Printf("aut Customer err -%v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		h.user.Customer = user
		w.Write([]byte("Привет, ты сегодня чудесно выглядишь "))

	} else if h.user.Role == "loader" {
		user, err := dbLoader.GetInfo(h.ctx, h.sql, h.user.Login)
		if err != nil {
			log.Printf("aut Loader err -%v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.user.Loader = user
		w.Write([]byte("Привет, не переживай все будет хорошо "))
	}
	token, err := interal.GenerateToken(h.user.Login, h.user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ctx = context.WithValue(h.ctx, "token", token)
	w.WriteHeader(http.StatusOK)
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
	token, err := interal.GenerateToken(h.user.Login, h.user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.ctx = context.WithValue(h.ctx, "token", token)

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
