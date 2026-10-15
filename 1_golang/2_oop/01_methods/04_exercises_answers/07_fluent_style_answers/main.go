// Упражнение 07: fluent-интерфейс (цепочка методов) (Решение).
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

func (q *Query) Table(t string) *Query {
	q.table = t
	return q
}

func (q *Query) Where(cond string) *Query {
	q.where = append(q.where, cond)
	return q
}

func (q *Query) Limit(n int) *Query {
	q.limit = n
	return q
}

func (q *Query) Build() string {
	s := "SELECT * FROM " + q.table
	if len(q.where) > 0 {
		s += " WHERE " + strings.Join(q.where, " AND ")
	}
	if q.limit > 0 {
		s += fmt.Sprintf(" LIMIT %d", q.limit)
	}
	return s
}

func main() {
	fmt.Println("--- Проверка 07_fluent_style (Решение) ---")

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
