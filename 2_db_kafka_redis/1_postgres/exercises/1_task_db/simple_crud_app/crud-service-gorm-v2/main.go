package main

import (
	"net/http"
)

func main() {
	db := InitDB()

	// Автоматическая миграция — создаст таблицы, если их нет
	db.AutoMigrate(&User{}, &Order{}, &Item{})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreateUserHandler(db)(w, r)
		case http.MethodGet:
			GetUserHandler(db)(w, r)
		case http.MethodPut:
			UpdateUserHandler(db)(w, r)
		case http.MethodDelete:
			DeleteUserHandler(db)(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8888", nil)
}
