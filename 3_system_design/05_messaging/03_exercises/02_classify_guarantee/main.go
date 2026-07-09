// Упражнение 02: классификация гарантии доставки.
//
// Задача: по описанию сценария вернуть уровень гарантии доставки:
//
//	"at-most-once"  — консьюмер сначала подтверждает (ack/commit), потом
//	                  обрабатывает: возможна потеря, дублей нет;
//	"at-least-once" — консьюмер сначала обрабатывает, потом подтверждает:
//	                  потерь нет, возможны дубли;
//	"exactly-once"  — at-least-once + дедупликация/идемпотентность: эффект
//	                  применяется ровно один раз.
//
// На вход даётся сценарий: порядок действий (ackFirst) и наличие дедупликации
// (dedup). Верните строку с уровнем гарантии.
//
//	ackFirst=true                 -> "at-most-once"
//	ackFirst=false, dedup=false   -> "at-least-once"
//	ackFirst=false, dedup=true    -> "exactly-once"
//
// Запуск: go run ./03_exercises/02_classify_guarantee/main.go
package main

import "fmt"

// classify возвращает уровень гарантии по параметрам сценария.
//
// ackFirst — консьюмер подтверждает ДО обработки (true) или ПОСЛЕ (false).
// dedup    — есть ли дедупликация/идемпотентность на стороне консьюмера.
func classify(ackFirst, dedup bool) string {
	// TODO: реализуйте классификацию по описанию выше.
	return ""
}

var _ = classify

func main() {
	fmt.Println("--- Проверка 02_classify_guarantee ---")

	cases := []struct {
		ackFirst bool
		dedup    bool
		want     string
	}{
		{true, false, "at-most-once"},
		{true, true, "at-most-once"}, // ack раньше обработки -> потеря важнее, дедуп не спасает
		{false, false, "at-least-once"},
		{false, true, "exactly-once"},
	}

	allOK := true
	for _, c := range cases {
		got := classify(c.ackFirst, c.dedup)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: classify(ackFirst=%v, dedup=%v) = %q, ожидалось %q\n",
				c.ackFirst, c.dedup, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
