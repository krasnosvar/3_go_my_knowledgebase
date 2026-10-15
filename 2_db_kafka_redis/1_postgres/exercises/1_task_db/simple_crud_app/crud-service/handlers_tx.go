package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateUserTx — создание пользователя в транзакции
func CreateUserTx(db *sql.DB, user User) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() // откат при ошибке

	var id int
	err = tx.QueryRow(
		"INSERT INTO users (order_count) VALUES ($1) RETURNING id",
		user.OrderCount,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateUserTx — обновление пользователя в транзакции
func UpdateUserTx(db *sql.DB, user User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec(
		"UPDATE users SET order_count=$1 WHERE id=$2",
		user.OrderCount, user.ID,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", user.ID)
	}

	return tx.Commit()
}

// DeleteUserTx — удаление пользователя с каскадом в транзакции
func DeleteUserTx(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// сначала удаляем items
	_, err = tx.Exec(`
        DELETE FROM items WHERE order_id IN (
            SELECT id FROM orders WHERE user_id=$1
        )`, id)
	if err != nil {
		return err
	}

	// затем удаляем orders
	_, err = tx.Exec("DELETE FROM orders WHERE user_id=$1", id)
	if err != nil {
		return err
	}

	// и только потом удаляем пользователя
	res, err := tx.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return tx.Commit()
}

// ───────────────────────────────────────────────
// Хендлеры для работы с транзакциями
// ───────────────────────────────────────────────

func CreateUserTxHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		id, err := CreateUserTx(db, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}

func UpdateUserTxHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := UpdateUserTx(db, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
	}
}

func DeleteUserTxHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		var id int
		fmt.Sscanf(idStr, "%d", &id)

		if err := DeleteUserTx(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
	}
}
