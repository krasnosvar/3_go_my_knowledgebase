// Пример 03: пробы liveness (/healthz) и readiness (/readyz).
//
// Поднимаем HTTP-сервер на 127.0.0.1:0 (случайный порт). /healthz
// всегда отвечает 200 (процесс жив). /readyz зависит от атомарного
// флага готовности: пока false -> 503, после переключения -> 200.
// Клиент дёргает оба эндпоинта до и после переключения readiness,
// затем сервер аккуратно останавливается — и программа завершается.
//
// Демонстрирует: разницу liveness/readiness, управление готовностью
// через atomic.Bool, сервер на :0 и graceful shutdown.
//
// Запуск: go run ./02_examples/03_health_readiness/main.go
package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	// ready — флаг готовности принимать трафик. На старте false
	// (например, идёт прогрев кэша / подключение к БД).
	var ready atomic.Bool

	mux := http.NewServeMux()

	// liveness: жив ли процесс. Не зависит от внешних зависимостей —
	// провал означал бы «перезапустите контейнер».
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	})

	// readiness: готов ли инстанс принимать запросы. Провал (503) ->
	// оркестратор убирает под из балансировки, но не перезапускает.
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		if ready.Load() {
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, "ready")
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = io.WriteString(w, "not ready")
	})

	// Слушаем на 127.0.0.1:0 — ядро выдаёт свободный порт.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	base := "http://" + ln.Addr().String()
	fmt.Printf("сервер слушает на %s\n", base)

	srv := &http.Server{Handler: mux}
	go func() { _ = srv.Serve(ln) }() // остановим через Shutdown

	client := &http.Client{Timeout: 5 * time.Second}

	// probe делает GET и печатает код статуса + тело.
	probe := func(path string) {
		resp, err := client.Get(base + path)
		if err != nil {
			fmt.Printf("%s -> ошибка: %v\n", path, err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s -> %d %q\n", path, resp.StatusCode, body)
	}

	fmt.Println("\n-- до готовности --")
	probe("/healthz") // 200: процесс жив
	probe("/readyz")  // 503: ещё не готов

	// Прогрев завершён — включаем готовность.
	ready.Store(true)

	fmt.Println("\n-- после ready.Store(true) --")
	probe("/healthz") // 200
	probe("/readyz")  // 200: теперь готов

	// Graceful shutdown — не оставляем висящий сервер.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)

	fmt.Println("\nГотово, программа завершается.")
}
