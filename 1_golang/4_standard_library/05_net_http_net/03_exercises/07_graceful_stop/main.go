// Упражнение 07: graceful shutdown HTTP-сервера.
//
// Задача: реализуйте run(srv, ln, stop), которая:
//   - запускает srv.Serve(ln) в отдельной горутине;
//   - ждёт сигнал из канала stop (в реальном сервисе это SIGINT/SIGTERM,
//     см. signal.NotifyContext);
//   - вызывает srv.Shutdown(ctx) с таймаутом 2s: сервер перестаёт
//     принимать новые соединения, но ДОЖИДАЕТСЯ уже начатых запросов;
//   - возвращает nil при чистой остановке. Ошибка Serve, равная
//     http.ErrServerClosed, — нормальный результат Shutdown, а не сбой.
//
// srv.Close() вместо Shutdown оборвал бы запросы на лету — тест это ловит.
//
// Запуск: go run ./03_exercises/07_graceful_stop/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

var _ = context.WithTimeout // держим импорты живыми в заготовке
var _ = errors.Is

// run обслуживает запросы до сигнала stop и корректно останавливает сервер.
func run(srv *http.Server, ln net.Listener, stop <-chan struct{}) error {
	// TODO: go Serve → <-stop → Shutdown(ctx с таймаутом) → дождаться Serve.
	go srv.Serve(ln)
	<-stop
	return srv.Close()
}

func main() {
	fmt.Println("--- Проверка 07_graceful_stop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	started := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		close(started)
		time.Sleep(200 * time.Millisecond) // «долгий» запрос на лету
		fmt.Fprint(w, "done")
	})
	srv := &http.Server{Handler: mux}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("FAIL: listen:", err)
		return
	}

	stop := make(chan struct{})
	runErr := make(chan error, 1)
	go func() { runErr <- run(srv, ln, stop) }()

	type result struct {
		status int
		err    error
	}
	reqDone := make(chan result, 1)
	go func() {
		c := &http.Client{Timeout: 2 * time.Second}
		resp, err := c.Get("http://" + ln.Addr().String() + "/slow")
		if err != nil {
			reqDone <- result{0, err}
			return
		}
		resp.Body.Close()
		reqDone <- result{resp.StatusCode, nil}
	}()

	select {
	case <-started:
	case <-time.After(2 * time.Second):
		fmt.Println("FAIL: запрос не дошёл до сервера")
		fmt.Println("FAIL")
		return
	}
	close(stop) // сигнал остановки, пока запрос ещё выполняется

	res := <-reqDone
	check(res.err == nil && res.status == http.StatusOK,
		fmt.Sprintf("запрос на лету должен завершиться 200, получено %d, %v", res.status, res.err))

	select {
	case err := <-runErr:
		check(err == nil, fmt.Sprintf("run должен вернуть nil, получено %v", err))
	case <-time.After(3 * time.Second):
		check(false, "run не вернулся после остановки")
	}

	_, err = net.DialTimeout("tcp", ln.Addr().String(), 200*time.Millisecond)
	check(err != nil, "после остановки новые соединения не должны приниматься")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
