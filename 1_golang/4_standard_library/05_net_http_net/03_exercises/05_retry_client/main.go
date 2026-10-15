// Упражнение 05: повтор запроса при ошибках 5xx.
//
// Задача: реализуйте getWithRetry(c, url, attempts, pause), которая:
//   - делает до attempts попыток GET;
//   - успех (статус < 500) → сразу возвращает статус и nil;
//   - сетевая ошибка или статус >= 500 → пауза pause и следующая попытка;
//   - тело каждого ответа закрывает (иначе соединения не переиспользуются);
//   - если попытки кончились → (последний статус или 0, ошибка).
//
// 4xx не повторяем: это ошибка клиента, повтор не поможет.
// В реальном коде пауза растёт экспоненциально и с jitter — см.
// 3_system_design/08_reliability_scalability.
//
// Запуск: go run ./03_exercises/05_retry_client/main.go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"time"
)

// getWithRetry выполняет GET с повторами на 5xx и сетевых ошибках.
func getWithRetry(c *http.Client, url string, attempts int, pause time.Duration) (int, error) {
	// TODO: цикл попыток, resp.Body.Close(), повтор только на err/5xx.
	return 0, nil
}

func main() {
	fmt.Println("--- Проверка 05_retry_client ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var calls atomic.Int32
	failFirst := int32(2)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := calls.Add(1)
		if r.URL.Path == "/bad-request" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if n <= failFirst {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()
	c := &http.Client{Timeout: time.Second}

	status, err := getWithRetry(c, ts.URL, 5, time.Millisecond)
	check(err == nil && status == 200, fmt.Sprintf("2 сбоя и успех: получено %d, %v", status, err))
	check(calls.Load() == 3, fmt.Sprintf("ожидалось 3 обращения к серверу, было %d", calls.Load()))

	calls.Store(0)
	status, err = getWithRetry(c, ts.URL, 2, time.Millisecond)
	check(err != nil && status == 503, fmt.Sprintf("попытки кончились: ожидалось 503 и ошибка, получено %d, %v", status, err))

	calls.Store(100) // дальше сервер отвечает 200 на всё, кроме /bad-request
	status, err = getWithRetry(c, ts.URL+"/bad-request", 5, time.Millisecond)
	check(err == nil && status == 400 && calls.Load() == 101, fmt.Sprintf("4xx не повторяется: статус %d, err %v, обращений %d", status, err, calls.Load()-100))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
