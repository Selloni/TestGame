package handler

import (
	"log"
	"net/http"
)

// todo: midlle, config
func Route() error {
	http.HandleFunc("/login", LoginHandle)
	log.Println("listen port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {

}
