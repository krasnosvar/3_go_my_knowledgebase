// Упражнение 03: Decorator — добавить поведение, не меняя исходный тип.
//
// Есть интерфейс Handler и базовая реализация GreetHandler. Менять
// GreetHandler НЕЛЬЗЯ. Нужно добавить префикс "[audit] " к результату
// через ДЕКОРАТОР — тип-обёртку, реализующую тот же интерфейс Handler.
//
// Реализуйте метод AuditHandler.Handle так, чтобы он:
//  1. вызвал вложенный h.next.Handle(input),
//  2. вернул результат с префиксом "[audit] ".
//
// Заготовка компилируется и печатает FAIL — доведите до PASS.
//
// Запуск: go run ./03_exercises/03_decorator/main.go
package main

import "fmt"

// Handler — абстракция с одним методом.
type Handler interface {
	Handle(input string) string
}

// GreetHandler — базовая реализация. НЕ ИЗМЕНЯЙТЕ её.
type GreetHandler struct{}

func (GreetHandler) Handle(input string) string { return "привет, " + input }

// AuditHandler — ДЕКОРАТОР вокруг Handler.
type AuditHandler struct {
	next Handler
}

// Handle должен добавить префикс "[audit] " к результату вложенного Handler.
func (h AuditHandler) Handle(input string) string {
	// TODO: вызовите h.next.Handle(input) и верните результат
	// с префиксом "[audit] ".
	return ""
}

func main() {
	fmt.Println("--- Проверка 03_decorator ---")

	// Оборачиваем базовый хэндлер декоратором, не меняя GreetHandler.
	var h Handler = AuditHandler{next: GreetHandler{}}

	allOK := true

	if got := h.Handle("мир"); got != "[audit] привет, мир" {
		allOK = false
		fmt.Printf("FAIL: Handle = %q, ожидалось %q\n", got, "[audit] привет, мир")
	}
	// Базовый хэндлер должен остаться неизменным.
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
