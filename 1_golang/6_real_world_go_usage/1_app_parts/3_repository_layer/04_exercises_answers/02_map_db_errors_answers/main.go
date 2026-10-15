// Упражнение 02: маппинг ошибок драйвера в доменные (Решение).
package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("duplicate")
)

// pgError имитирует ошибку драйвера PostgreSQL с SQLSTATE-кодом.
type pgError struct{ Code string }

func (e *pgError) Error() string { return "pg error " + e.Code }

func mapErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	var pe *pgError
	if errors.As(err, &pe) && pe.Code == "23505" {
		return ErrDuplicate
	}
	return fmt.Errorf("repo: %w", err)
}

func main() {
	fmt.Println("--- Проверка 02_map_db_errors (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(mapErr(nil) == nil, "nil → nil")
	check(errors.Is(mapErr(fmt.Errorf("scan: %w", sql.ErrNoRows)), ErrNotFound), "обёрнутый sql.ErrNoRows → ErrNotFound")
	check(errors.Is(mapErr(fmt.Errorf("insert: %w", &pgError{Code: "23505"})), ErrDuplicate), "код 23505 → ErrDuplicate")
	other := &pgError{Code: "40001"}
	got := mapErr(other)
	check(!errors.Is(got, ErrDuplicate) && errors.Is(got, other) && got != error(other), "прочие ошибки обёрнуты с префиксом repo")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
