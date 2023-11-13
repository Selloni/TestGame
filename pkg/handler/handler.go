package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// todo: midlle, config
func Route() error {
	// open request
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/register", registerHandle)
	http.HandleFunc("/tasks", tasksHandler)

	log.Println("listen port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	role := r.FormValue("role")
	fmt.Println(role)
	// ищем пользователя в бд и отправляем токен
	http.Error(w, "пользователь или пароль не верен", http.StatusUnauthorized)
}

func registerHandle(w http.ResponseWriter, r *http.Request) {
	// поиск по логину, если нет добавляем в бд
	if true {
		http.Error(w, "пользователь уже существует", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
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
