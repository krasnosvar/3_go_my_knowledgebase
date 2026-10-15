// Упражнение 08: graceful shutdown пула worker'ов.
//
// Задача: Server — пул worker'ов с корректной остановкой.
//   - newServer(workers, handle) создаёт буферизованный канал jobs и
//     запускает workers горутин: for j := range s.jobs { handle(j) };
//   - Submit(v) ставит задание в очередь (true) — до вызова Shutdown;
//   - Shutdown(ctx) закрывает jobs (новых заданий больше нет), ждёт, пока
//     worker'ы доработают очередь, и возвращает nil; если ctx истёк
//     раньше — возвращает ctx.Err().
//
// Ожидание wg.Wait() с таймаутом: запустите его в горутине, которая
// закроет канал done, и сделайте select по done и ctx.Done().
//
// Запуск: go run ./03_exercises/08_graceful_shutdown/main.go
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Server — пул worker'ов с graceful shutdown.
type Server struct {
	jobs chan int
	wg   sync.WaitGroup
}

func newServer(workers int, handle func(int)) *Server {
	// TODO: s := &Server{jobs: make(chan int, 100)}; запустить workers горутин.
	return &Server{}
}

// Submit ставит задание в очередь.
func (s *Server) Submit(v int) bool {
	// TODO
	return false
}

// Shutdown останавливает приём и ждёт завершения или истечения ctx.
func (s *Server) Shutdown(ctx context.Context) error {
	// TODO: close(s.jobs); ждать s.wg с учётом ctx.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_graceful_shutdown ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var processed atomic.Int64
	s := newServer(3, func(int) {
		time.Sleep(time.Millisecond)
		processed.Add(1)
	})
	for i := 0; i < 10; i++ {
		check(s.Submit(i), fmt.Sprintf("Submit(%d) до Shutdown должен пройти", i))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	err := s.Shutdown(ctx)
	cancel()
	check(err == nil, fmt.Sprintf("Shutdown вернул %v, ожидался nil", err))
	check(processed.Load() == 10, fmt.Sprintf("обработано %d из 10 — очередь не доработана", processed.Load()))

	release := make(chan struct{})
	stuck := newServer(1, func(int) { <-release })
	stuck.Submit(1)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 20*time.Millisecond)
	err = stuck.Shutdown(ctx2)
	cancel2()
	check(errors.Is(err, context.DeadlineExceeded), fmt.Sprintf("зависший worker: %v, ожидалось context.DeadlineExceeded", err))
	close(release)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
