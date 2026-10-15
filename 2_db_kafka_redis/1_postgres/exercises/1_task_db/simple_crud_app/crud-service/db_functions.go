// В handlers.go появятся функции вроде:
// CreateUser(db *sql.DB, user User) (int, error) → вставка нового пользователя.
// GetUser(db *sql.DB, id int) (User, error) → выборка пользователя по ID.
// UpdateUser(db *sql.DB, user User) error → обновление данных.
// DeleteUser(db *sql.DB, id int) error → удаление.
// Эти функции будут использовать db.QueryRow, db.Exec, db.Prepare.
// Потом мы свяжем их с HTTP‑эндпоинтами (/users, /orders, /items), чтобы можно было дергать их через REST API.

package main

import (
	"database/sql"
	"fmt"
)

// ───────────────────────────────────────────────
// Шпаргалка по database/sql
// ───────────────────────────────────────────────
//
// db.QueryRow → одна строка
//   Пример: SELECT ... WHERE id=$1
//   Используется с .Scan(&var1, &var2)
//   Ошибка sql.ErrNoRows если ничего не найдено
//
// db.Query → несколько строк
//   Пример: SELECT ... FROM table
//   Возвращает rows, перебираем через rows.Next()
//   Обязательно rows.Close() после использования
//
// db.Exec → запрос без возврата строк
//   Пример: INSERT, UPDATE, DELETE
//   Возвращает результат, можно проверить RowsAffected()
//   Удобно для массовых операций
//
// db.Prepare → подготовка запроса
//   Пример: stmt, _ := db.Prepare("INSERT ...")
//   Используется многократно с stmt.Exec()
//   После работы stmt.Close()
//
// db.QueryRow с RETURNING → получить значение сразу
//   Пример: INSERT INTO ... RETURNING id
//   Используется с .Scan(&id)
//   В PostgreSQL это замена LastInsertId()
// ───────────────────────────────────────────────

func CreateUser(db *sql.DB, user User) (int, error) {
	// 	Не нужно указывать id при вставке
	// В PostgreSQL обычно поле id — это SERIAL или BIGSERIAL. Оно генерируется автоматически.
	// Поэтому в INSERT лучше писать только order_count.
	var id int

	err := db.QueryRow(
		// QueryRow — выполняет запрос, который возвращает одну строку.
		// .Scan(...) — извлекает значения из этой строки в переменные.
		// Если строка не найдена, вернётся ошибка sql.ErrNoRows
		"INSERT INTO users (order_count) VALUES ($1) RETURNING id",
		user.OrderCount,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUser(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow(
		"SELECT id, order_count FROM users WHERE id=$1",
		id,
	).Scan(&user.ID, &user.OrderCount)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user User) error {
	// Exec — выполняет запрос, который не возвращает строки (INSERT, UPDATE, DELETE).
	// RowsAffected() — показывает, сколько строк было изменено.
	// Используется для операций изменения данных.
	res, err := db.Exec(
		"UPDATE users SET order_count=$1 WHERE id=$2",
		user.OrderCount, user.ID,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", user.ID)
	}

	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	res, err := db.Exec(
		"DELETE FROM users WHERE id=$1",
		id,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}

// TOP-10 TASK
// SELECT u.id, u.order_count, SUM(i.price) AS total_spent   -- выбираем id пользователя, количество заказов и сумму цен его товаров
// FROM users u                                             -- основная таблица: пользователи
// JOIN orders o ON u.id = o.user_id                        -- соединяем с таблицей заказов по user_id
// JOIN items i ON o.id = i.order_id                        -- соединяем с таблицей товаров по order_id
// WHERE o.created_at >= NOW() - INTERVAL '30 days'         -- фильтруем: только заказы за последние 30 дней
// GROUP BY u.id, u.order_count                             -- группируем по пользователю (id и количество заказов)
// ORDER BY total_spent DESC                                -- сортируем по сумме покупок, от большего к меньшему
// LIMIT 10;                                                -- берём только первых 10 пользователей

type UserStats struct {
	ID         int
	OrderCount int
	TotalSpent float64
}

func GetTopUsers(db *sql.DB) ([]UserStats, error) {
	// Query — выполняет запрос, который возвращает несколько строк.
	// rows.Next() — итерация по результатам.
	// .Scan(...) — извлекает значения из текущей строки.
	// defer rows.Close() — обязательно закрываем курсор.
	rows, err := db.Query(`
        SELECT u.id, u.order_count, SUM(i.price) AS total_spent
        FROM users u
        JOIN orders o ON u.id = o.user_id
        JOIN items i ON o.id = i.order_id
        WHERE o.created_at >= NOW() - INTERVAL '30 days'
        GROUP BY u.id, u.order_count
        ORDER BY total_spent DESC
        LIMIT 10;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []UserStats
	for rows.Next() {
		var us UserStats
		// обрабатываем каждую строку
		if err := rows.Scan(&us.ID, &us.OrderCount, &us.TotalSpent); err != nil {
			return nil, err
		}
		result = append(result, us)
	}
	return result, nil
}
