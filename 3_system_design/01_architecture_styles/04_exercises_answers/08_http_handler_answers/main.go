// Упражнение 08: HTTP-хэндлер с правильным статусом и заголовком (Решение).
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		// Content-Type по умолчанию для короткого текста нас устроит;
		// главное — выставить статус до записи тела.
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "missing name")
		return
	}
	// Заголовки выставляем ДО WriteHeader.
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, "hello, "+name)
}

func main() {
	fmt.Println("--- Проверка 08_http_handler (Решение) ---")

	srv := httptest.NewServer(http.HandlerFunc(greetHandler))
	defer srv.Close()

	allOK := true

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
