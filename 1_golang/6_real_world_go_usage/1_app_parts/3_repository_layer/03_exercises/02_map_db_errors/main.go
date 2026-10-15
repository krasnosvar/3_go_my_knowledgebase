// Упражнение 02: маппинг ошибок драйвера в доменные.
//
// Задача: реализуйте mapErr(err error) error для слоя репозитория:
//   - nil → nil;
//   - sql.ErrNoRows (в т.ч. обёрнутая) → ErrNotFound;
//   - ошибка драйвера с кодом 23505 (unique_violation в PostgreSQL) → ErrDuplicate;
//     код достаётся через errors.As в *pgError;
//   - остальное → fmt.Errorf("repo: %w", err).
//
// Запуск: go run ./03_exercises/02_map_db_errors
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

var _ = sql.ErrNoRows

func mapErr(err error) error {
	// TODO: errors.Is для sql.ErrNoRows, errors.As для *pgError.
	return err
}

func main() {
	fmt.Println("--- Проверка 02_map_db_errors ---")

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
