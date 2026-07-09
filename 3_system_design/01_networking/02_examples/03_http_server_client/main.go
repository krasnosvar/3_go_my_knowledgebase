// Пример 03: HTTP-сервер и клиент на net/http.
//
// Поднимаем HTTP-сервер через httptest.NewServer (сам выбирает
// случайный порт на 127.0.0.1), делаем GET-запрос клиентом, печатаем
// статус, интересующие заголовки и тело — и программа завершается.
//
// Демонстрирует: http.HandlerFunc, httptest.NewServer, http.Client
// с таймаутом, чтение resp.StatusCode / resp.Header / resp.Body.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	// Хэндлер отвечает JSON'ом и задаёт заголовок Content-Type.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Сервер получил запрос: %s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"message":"hello http"}`)
	})

	// httptest.NewServer поднимает реальный HTTP-сервер на 127.0.0.1
	// со случайным портом. Close() гарантированно останавливает его.
	srv := httptest.NewServer(handler)
	defer srv.Close()
	fmt.Printf("HTTP-сервер слушает на %s\n", srv.URL)

	// Клиент с таймаутом — чтобы запрос никогда не завис навсегда.
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(srv.URL + "/hello")
	if err != nil {
		log.Fatalf("GET: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read body: %v", err)
	}

	fmt.Printf("Статус: %s (%d)\n", resp.Status, resp.StatusCode)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("Тело: %s\n", body)
	fmt.Println("Готово, программа завершается.")
}
