// Упражнение 04: HTTP-хэндлер с правильным статусом и заголовком.
//
// Задача: реализовать greetHandler — HTTP-обработчик, который:
//   - на путь с query-параметром ?name=... отвечает 200 OK,
//     заголовком Content-Type: text/plain и телом "hello, <name>";
//   - если параметр name пуст или отсутствует — отвечает 400
//     Bad Request и телом "missing name".
//
// Сервер поднимается через httptest и делает два запроса для проверки.
//
// Запуск: go run ./03_exercises/04_http_handler/main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

// greetHandler — реализуйте согласно условию.
func greetHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:
	//  1) name := r.URL.Query().Get("name");
	//  2) если name == "" -> w.WriteHeader(http.StatusBadRequest) и тело "missing name";
	//  3) иначе -> Content-Type text/plain, статус 200, тело "hello, "+name.
	//
	// Внимание: заголовки (w.Header().Set) нужно выставлять ДО WriteHeader.
}

func main() {
	fmt.Println("--- Проверка 04_http_handler ---")

	srv := httptest.NewServer(http.HandlerFunc(greetHandler))
	defer srv.Close()

	allOK := true

	// Случай 1: с именем -> 200 и "hello, Alice".
	resp, err := http.Get(srv.URL + "/?name=Alice")
	if err != nil {
		log.Fatalf("GET: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK || string(body) != "hello, Alice" {
		allOK = false
		fmt.Printf("FAIL: с именем: статус=%d, тело=%q\n", resp.StatusCode, body)
	}
	if ct := resp.Header.Get("Content-Type"); ct != "text/plain" {
		allOK = false
		fmt.Printf("FAIL: Content-Type=%q, ожидался text/plain\n", ct)
	}

	// Случай 2: без имени -> 400 и "missing name".
	resp2, err := http.Get(srv.URL + "/")
	if err != nil {
		log.Fatalf("GET: %v", err)
	}
	body2, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()
	if resp2.StatusCode != http.StatusBadRequest || string(body2) != "missing name" {
		allOK = false
		fmt.Printf("FAIL: без имени: статус=%d, тело=%q\n", resp2.StatusCode, body2)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
