// Пример 10: graceful shutdown HTTP-сервера.
//
// Поднимаем http.Server на 127.0.0.1:0 (порт выбирает ОС). Отдельная
// горутина делает запрос к серверу. Затем инициируем корректное
// завершение через Server.Shutdown(ctx): сервер перестаёт принимать
// новые соединения и дожидается завершения активных запросов (дренаж),
// после чего программа выходит.
//
// Никаких висящих серверов: порт :0, обязательный Shutdown с таймаутом,
// программа гарантированно завершается.
//
// В реальном сервисе сигналом к остановке служит SIGTERM; здесь показан
// идиоматичный приём через signal.NotifyContext, но триггерим отмену
// сами, чтобы пример был детерминированным и завершался.
//
// Запуск: go run ./02_examples/10_graceful_shutdown/main.go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("--- graceful shutdown ---")

	// Хэндлер с небольшой «работой», чтобы показать дренаж активного запроса.
	mux := http.NewServeMux()
	mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("сервер: начал обработку /work")
		time.Sleep(100 * time.Millisecond) // имитация работы
		fmt.Println("сервер: закончил обработку /work")
		_, _ = io.WriteString(w, "done")
	})

	// Слушатель на 127.0.0.1:0 — ОС выберет свободный порт.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	addr := ln.Addr().String()
	fmt.Printf("сервер слушает на %s\n", addr)

	srv := &http.Server{Handler: mux}

	// context, который отменяется по SIGINT/SIGTERM (идиома graceful shutdown).
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	// Запускаем сервер в фоне.
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Serve вернёт ErrServerClosed после Shutdown — это штатно.
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Printf("serve: %v", err)
		}
	}()

	// Клиент в отдельной горутине шлёт запрос.
	wg.Add(1)
	go func() {
		defer wg.Done()
		client := &http.Client{Timeout: 2 * time.Second}
		resp, err := client.Get("http://" + addr + "/work")
		if err != nil {
			fmt.Printf("клиент: ошибка запроса: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("клиент: получил ответ %q (статус %d)\n", string(body), resp.StatusCode)
	}()

	// Даём запросу стартовать, затем триггерим остановку.
	// В реальности здесь просто ждали бы <-ctx.Done() (сигнал SIGTERM).
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("инициируем shutdown (эквивалент SIGTERM)")
		stop() // отменяем ctx, как если бы пришёл сигнал
	}()

	// Ждём сигнала к завершению.
	<-ctx.Done()

	// Дренаж с крайним сроком: активным запросам даём доиграть <=5с.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("сервер: начинаю graceful shutdown (дренаж соединений)")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: не уложились в таймаут: %v", err)
	} else {
		fmt.Println("сервер: все активные запросы завершены, listener закрыт")
	}

	wg.Wait() // дождаться и сервера, и клиента
	fmt.Println("Готово, программа завершается.")
}
