// Упражнение 03: Decorator — добавить поведение, не меняя исходный тип (Решение).
package main

import "fmt"

type Handler interface {
	Handle(input string) string
}

type GreetHandler struct{}

func (GreetHandler) Handle(input string) string { return "привет, " + input }

type AuditHandler struct {
	next Handler
}

func (h AuditHandler) Handle(input string) string {
	return "[audit] " + h.next.Handle(input)
}

func main() {
	fmt.Println("--- Проверка 03_decorator (Решение) ---")

	var h Handler = AuditHandler{next: GreetHandler{}}

	allOK := true

	if got := h.Handle("мир"); got != "[audit] привет, мир" {
		allOK = false
		fmt.Printf("FAIL: Handle = %q, ожидалось %q\n", got, "[audit] привет, мир")
	}
	if got := (GreetHandler{}).Handle("мир"); got != "привет, мир" {
		allOK = false
		fmt.Printf("FAIL: базовый GreetHandler изменён: %q\n", got)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
