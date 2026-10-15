// Упражнение 08: сервис здоровья (grpc.health.v1).
//
// Задача: реализуйте healthServer:
//   - SetServingStatus(service string, s servingStatus);
//   - Check(service string) (servingStatus, error):
//     зарегистрированный сервис → его статус;
//     неизвестный → (ServiceUnknown, errNotFound);
//     пустое имя "" означает «сервер в целом»: SERVING, только если ВСЕ
//     зарегистрированные сервисы SERVING (и хотя бы один зарегистрирован),
//     иначе NOT_SERVING.
//
// Методы вызываются конкурентно — защитите map мьютексом.
//
// Запуск: go run ./03_exercises/08_health_rpc
package main

import (
	"errors"
	"fmt"
	"sync"
)

type servingStatus int

const (
	ServiceUnknown servingStatus = iota
	Serving
	NotServing
)

var errNotFound = errors.New("unknown service")

type healthServer struct {
	mu       sync.Mutex
	statuses map[string]servingStatus
}

func (h *healthServer) SetServingStatus(service string, s servingStatus) {
	// TODO: запишите статус под мьютексом.
}

func (h *healthServer) Check(service string) (servingStatus, error) {
	// TODO: обработайте "" отдельно, неизвестный → errNotFound.
	return ServiceUnknown, nil
}

func main() {
	fmt.Println("--- Проверка 08_health_rpc ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	h := &healthServer{statuses: map[string]servingStatus{}}
	s, _ := h.Check("")
	check(s == NotServing, "пустой сервер → NOT_SERVING")

	var wg sync.WaitGroup
	for _, name := range []string{"orders", "users"} {
		wg.Add(1)
		go func(n string) { defer wg.Done(); h.SetServingStatus(n, Serving) }(name)
	}
	wg.Wait()

	s, err := h.Check("orders")
	check(err == nil && s == Serving, "orders → SERVING")
	s, _ = h.Check("")
	check(s == Serving, "все SERVING → сервер SERVING")
	h.SetServingStatus("users", NotServing)
	s, _ = h.Check("")
	check(s == NotServing, "один NOT_SERVING → сервер NOT_SERVING")
	s, err = h.Check("billing")
	check(s == ServiceUnknown && errors.Is(err, errNotFound), "неизвестный → SERVICE_UNKNOWN + ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
