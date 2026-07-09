// Упражнение 01: идемпотентный обработчик (дедуп по id).
//
// Задача: реализуйте handler, который принимает сообщения (возможны дубли по
// id, как при at-least-once) и применяет их эффект РОВНО ОДИН РАЗ на id.
//
// handler.Handle(id, amount):
//   - если id ещё не обрабатывался -> прибавить amount к сумме, вернуть true
//     (эффект применён);
//   - если id уже обрабатывался (дубль) -> НИЧЕГО не менять, вернуть false.
//
// handler.Total() -> текущая сумма применённых amount.
//
// Подсказка: храните обработанные id в map[string]bool.
//
// Запуск: go run ./03_exercises/01_idempotent_handler/main.go
package main

import "fmt"

// idempotentHandler дедуплицирует сообщения по id.
type idempotentHandler struct {
	seen  map[string]bool
	total int
}

func newHandler() *idempotentHandler {
	// TODO: инициализируйте поля (не забудьте про map).
	return &idempotentHandler{}
}

// Handle применяет эффект один раз на id. Возвращает true, если эффект
// применён впервые, и false для дубля.
func (h *idempotentHandler) Handle(id string, amount int) bool {
	// TODO: реализуйте дедупликацию по id.
	return false
}

// Total возвращает сумму применённых amount.
func (h *idempotentHandler) Total() int {
	// TODO: верните накопленную сумму.
	return 0
}

// Держим тип «живым» на случай, если решение вынесет его в пакет.
var _ = newHandler

func main() {
	fmt.Println("--- Проверка 01_idempotent_handler ---")

	h := newHandler()

	type ev struct {
		id     string
		amount int
		want   bool // ожидаемый результат Handle (применён впервые?)
	}
	events := []ev{
		{"a", 10, true},
		{"b", 5, true},
		{"a", 10, false}, // дубль
		{"c", 1, true},
		{"b", 5, false},  // дубль
		{"a", 10, false}, // дубль
	}

	allOK := true
	for _, e := range events {
		got := h.Handle(e.id, e.amount)
		if got != e.want {
			allOK = false
			fmt.Printf("FAIL: Handle(%q,%d) = %v, ожидалось %v\n", e.id, e.amount, got, e.want)
		}
	}

	const wantTotal = 16 // 10 + 5 + 1, дубли не учитываются
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
