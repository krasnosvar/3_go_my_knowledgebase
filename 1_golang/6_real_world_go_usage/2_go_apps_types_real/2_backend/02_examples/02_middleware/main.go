// Пример: logging middleware вокруг mux. Сервер поднимается через httptest,
// выполняются два запроса, в лог (stdout) пишутся метод, путь и время
// обработки; программа завершается сама.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func logging(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Println(r.Method, r.URL.Path, time.Since(start).Round(time.Millisecond))
	})
}

func main() {
	logger := log.New(os.Stdout, "[http] ", 0)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })

	srv := httptest.NewServer(logging(logger, mux))
	defer srv.Close()

	for _, p := range []string{"/", "/users"} {
		resp, err := http.Get(srv.URL + p)
		if err != nil {
			fmt.Println("request failed:", err)
			return
		}
		resp.Body.Close()
	}
}
