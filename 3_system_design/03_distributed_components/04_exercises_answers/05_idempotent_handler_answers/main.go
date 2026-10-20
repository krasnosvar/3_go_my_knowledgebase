// Упражнение 05: идемпотентный обработчик (Решение).
package main

import "fmt"

type idempotentHandler struct {
	seen  map[string]bool
	total int
}

func newHandler() *idempotentHandler {
	return &idempotentHandler{seen: make(map[string]bool)}
}

func (h *idempotentHandler) Handle(id string, amount int) bool {
	if h.seen[id] {
		return false // дубль -> эффект не применяем
	}
	h.seen[id] = true
	h.total += amount
	return true
}

func (h *idempotentHandler) Total() int {
	return h.total
}

var _ = newHandler

func main() {
	fmt.Println("--- Проверка 05_idempotent_handler (Решение) ---")

	h := newHandler()

	type ev struct {
		id     string
		amount int
		want   bool
	}
	events := []ev{
		{"a", 10, true},
		{"b", 5, true},
		{"a", 10, false},
		{"c", 1, true},
		{"b", 5, false},
		{"a", 10, false},
	}

	allOK := true
	for _, e := range events {
		got := h.Handle(e.id, e.amount)
		if got != e.want {
			allOK = false
			fmt.Printf("FAIL: Handle(%q,%d) = %v, ожидалось %v\n", e.id, e.amount, got, e.want)
		}
	}

	const wantTotal = 16
	if h.Total() != wantTotal {
		allOK = false
		fmt.Printf("FAIL: Total() = %d, ожидалось %d\n", h.Total(), wantTotal)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
