// Упражнение 04: HTTP-клиент с таймаутом (Решение).
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

// newClient создаёт HTTP-клиент с общим таймаутом запроса.
func newClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}

// fetch выполняет GET и возвращает тело ответа.
func fetch(c *http.Client, url string) (string, error) {
	resp, err := c.Get(url)
	if err != nil {
		return "", fmt.Errorf("get %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get %s: unexpected status %d", url, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}
	return string(body), nil
}

func main() {
	fmt.Println("--- Проверка 04_client_timeout (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "ok") })
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(300 * time.Millisecond):
			fmt.Fprint(w, "late")
		case <-r.Context().Done(): // клиент ушёл — не держим горутину
		}
	})
	mux.HandleFunc("/missing", http.NotFound)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	c := newClient(50 * time.Millisecond)
	check(c.Timeout == 50*time.Millisecond, fmt.Sprintf("Timeout клиента = %v, ожидалось 50ms", c.Timeout))

	body, err := fetch(c, ts.URL+"/fast")
	check(err == nil && body == "ok", fmt.Sprintf("/fast: получено %q, %v", body, err))

	start := time.Now()
	_, err = fetch(c, ts.URL+"/slow")
	check(err != nil, "/slow должен завершиться ошибкой таймаута")
	check(time.Since(start) < 250*time.Millisecond, "таймаут должен сработать раньше ответа медленного сервера")

	_, err = fetch(c, ts.URL+"/missing")
	check(err != nil, "статус 404 должен давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
