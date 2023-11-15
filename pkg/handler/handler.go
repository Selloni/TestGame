package handler

import (
	"WB/interal"
	"WB/interal/customer"
	dbCustomer "WB/interal/customer/db"
	"WB/interal/loader"
	dbLoader "WB/interal/loader/db"
	"WB/interal/posgresql"
	dbTask "WB/interal/task/db"
	"context"
	"fmt"
	"log"
	"net/http"
)

// todo: midlle, config

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
	} else {
		http.Error(w, "не определили роль", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
	var x any = h.ctx.Value("token")
	if x == nil {
		if err := h.taskPublic(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
	isValid, _, role := interal.ValidateToken(x.(string))
	if !isValid {
		http.Error(w, "не валидный токен", http.StatusUnauthorized)
		return
	}
	if role == "customer" {
		if err := h.taskCustomer(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else if role == "loader" {
		fmt.Println("logika")
	}
}

func (h *handler) taskPublic(w http.ResponseWriter) error {
	mm, err := dbTask.CreateTask(h.ctx, h.sql)
	dbTask.CreateTask(h.ctx, h.sql)
	if err != nil {
		return err
	}
	fmt.Fprint(w, "Создали новые задания\n")
	for k, v := range mm {
		fmt.Fprint(w, fmt.Sprintf("товар %s - вес %d\n", k, v))
	}
	return nil
}

func (h *handler) taskCustomer(w http.ResponseWriter) error {
	log.Printf("Get task for customer")
	allTask, err := dbCustomer.GetAllTask(h.ctx, h.sql)
	if err != nil {
		log.Printf("err %v", err)
		return err
	}
	for _, task := range allTask {
		fmt.Fprint(w, fmt.Sprintf("id - %d, name - %s, weight - %d",
			task.Id, task.Name, task.Weight), "\n")
	}
	return nil

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
