// Упражнение 08: graceful shutdown пула worker'ов (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	jobs chan int
	wg   sync.WaitGroup
}

func newServer(workers int, handle func(int)) *Server {
	s := &Server{jobs: make(chan int, 100)}
	for i := 0; i < workers; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for j := range s.jobs {
				handle(j)
			}
		}()
	}
	return s
}

func (s *Server) Submit(v int) bool {
	s.jobs <- v
	return true
}

func (s *Server) Shutdown(ctx context.Context) error {
	close(s.jobs)
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	fmt.Println("--- Проверка 08_graceful_shutdown (Решение) ---")

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
