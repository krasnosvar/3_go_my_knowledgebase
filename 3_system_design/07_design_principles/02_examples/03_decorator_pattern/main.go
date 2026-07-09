// Пример 03: паттерн Decorator в идиоматичном Go.
//
// Decorator — обёртка вокруг объекта, добавляющая поведение БЕЗ изменения
// исходного типа. В Go это тип-обёртка вокруг интерфейса (та же сигнатура,
// внутри вызывается вложенная реализация плюс дополнительное поведение).
//
// Показаны два классических применения:
//
//  1. Декоратор над простым интерфейсом Service: добавляем логирование.
//  2. HTTP middleware над http.Handler — тот же приём, что применяет
//     вся экосистема net/http (логирование поверх бизнес-хэндлера).
//
// Декораторы можно вкладывать друг в друга (стек обёрток). Программа
// детерминированно демонстрирует оба и завершается.
//
// Запуск: go run ./02_examples/03_decorator_pattern/main.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// ---------------------------------------------------------------------------
// 1. Декоратор над обычным интерфейсом.
// ---------------------------------------------------------------------------

// Service — абстракция с одним методом.
type Service interface {
	Do(x int) int
}

// SquareService — базовая реализация (возводит в квадрат).
type SquareService struct{}

func (SquareService) Do(x int) int { return x * x }

// LoggingService — ДЕКОРАТОР: оборачивает любой Service и добавляет
// логирование, не трогая исходный тип. Сам реализует Service, поэтому
// подставляется всюду, где ждут Service (и может оборачивать другой декоратор).
type LoggingService struct {
	next Service
}

func (l LoggingService) Do(x int) int {
	res := l.next.Do(x) // делегируем вложенной реализации
	fmt.Printf("[log] Do(%d) = %d\n", x, res)
	return res
}

// CountingService — ещё один декоратор: считает вызовы.
type CountingService struct {
	next  Service
	calls *int
}

func (c CountingService) Do(x int) int {
	*c.calls++
	return c.next.Do(x)
}

// ---------------------------------------------------------------------------
// 2. HTTP middleware — тот же паттерн поверх http.Handler.
// ---------------------------------------------------------------------------

// withLogging оборачивает http.Handler, логируя каждый запрос до вызова next.
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[http] %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("=== Декоратор над интерфейсом Service ===")

	calls := 0
	// Стек декораторов: логирование -> счётчик -> базовая реализация.
	var svc Service = SquareService{}
	svc = CountingService{next: svc, calls: &calls}
	svc = LoggingService{next: svc}

	fmt.Println("результат:", svc.Do(3))
	fmt.Println("результат:", svc.Do(5))
	fmt.Printf("всего вызовов базового сервиса: %d\n", calls)

	fmt.Println()
	fmt.Println("=== Декоратор (middleware) над http.Handler ===")

	// Бизнес-хэндлер ничего не знает о логировании.
	business := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	// Оборачиваем его декоратором withLogging.
	srv := httptest.NewServer(withLogging(business))
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/hello")
	if err != nil {
		fmt.Println("ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("ответ сервера: %q (status %d)\n", string(body), resp.StatusCode)

	fmt.Println()
	fmt.Println("Готово, программа завершается.")
}
