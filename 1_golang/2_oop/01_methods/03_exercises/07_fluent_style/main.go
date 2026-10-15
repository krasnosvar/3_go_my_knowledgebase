// Упражнение 07: fluent-интерфейс (цепочка методов).
//
// Задача: реализуйте построитель SQL-запроса в fluent-стиле. Каждый метод
// меняет билдер и возвращает его же (*Query), чтобы вызовы можно было
// сцеплять:
//   - Table(t) — имя таблицы;
//   - Where(cond) — добавляет условие (условия объединяются через " AND ");
//   - Limit(n) — ограничение (0 — без LIMIT);
//   - Build() — итоговая строка:
//     "SELECT * FROM users WHERE age > 18 AND active LIMIT 10".
//
// Без условий — без WHERE, без лимита — без LIMIT.
//
// Запуск: go run ./03_exercises/07_fluent_style/main.go
package main

import (
	"fmt"
	"strings"
)

// Query — построитель запроса.
type Query struct {
	table string
	where []string
	limit int
}

// Держим импорт strings живым в заготовке (в решении — strings.Join).
var _ = strings.Join

// Table задаёт таблицу.
func (q *Query) Table(t string) *Query {
	// TODO: сохраните t и верните q.
	return q
}

// Where добавляет условие.
func (q *Query) Where(cond string) *Query {
	// TODO: добавьте cond в q.where и верните q.
	return q
}

// Limit задаёт лимит.
func (q *Query) Limit(n int) *Query {
	// TODO: сохраните n и верните q.
	return q
}

// Build собирает строку запроса.
func (q *Query) Build() string {
	// TODO: "SELECT * FROM <table>[ WHERE a AND b][ LIMIT n]".
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_fluent_style ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := (&Query{}).Table("users").Where("age > 18").Where("active").Limit(10).Build()
	want := "SELECT * FROM users WHERE age > 18 AND active LIMIT 10"
	check(got == want, fmt.Sprintf("полный запрос: ожидалось %q, получено %q", want, got))

	got = (&Query{}).Table("orders").Build()
	check(got == "SELECT * FROM orders", fmt.Sprintf("без условий и лимита: получено %q", got))

	got = (&Query{}).Table("logs").Limit(5).Build()
	check(got == "SELECT * FROM logs LIMIT 5", fmt.Sprintf("только лимит: получено %q", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
