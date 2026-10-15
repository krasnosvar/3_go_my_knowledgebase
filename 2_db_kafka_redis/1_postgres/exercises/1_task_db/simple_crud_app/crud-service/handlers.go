package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

// Handler для создания пользователя
// пример curl-запроса чтобы вызвать хендлер
// # Отправляем POST-запрос на эндпоинт /users
// curl -X POST http://localhost:8080/users \
//      -H "Content-Type: application/json" \
//      -d '{"order_count": 5}'

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		// читаем JSON из тела запроса
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		id, err := CreateUser(db, user) // вызываем нашу функцию из models.go
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// возвращаем ID нового пользователя в JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

// Handler для получения пользователя по ID
// curl "http://localhost:8080/users?id=1"

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// достаём параметр id из query string
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}

		user, err := GetUser(db, id) // вызываем функцию из models.go
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "user not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user) // возвращаем JSON
	}
}

// Хендлер для обновления пользователя (PUT /users)
func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		err := UpdateUser(db, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
	}
}

// Хендлер для удаления пользователя (DELETE /users)
func DeleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}

		err = DeleteUser(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
	}
}
