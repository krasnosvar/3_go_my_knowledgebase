// Упражнение 07: happens-before через закрытие канала (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Mailbox struct {
	data  string
	ready chan struct{}
}

func newMailbox() *Mailbox { return &Mailbox{ready: make(chan struct{})} }

func (m *Mailbox) Put(s string) {
	m.data = s
	close(m.ready)
}

func (m *Mailbox) Take(timeout time.Duration) (string, bool) {
	select {
	case <-m.ready:
		return m.data, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	fmt.Println("--- Проверка 07_ordering (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	m := newMailbox()
	go func() {
		time.Sleep(5 * time.Millisecond)
		m.Put("payload")
	}()
	v, ok := m.Take(300 * time.Millisecond)
	check(ok && v == "payload", fmt.Sprintf("Take = (%q, %t), ожидалось (\"payload\", true)", v, ok))

	// несколько читателей видят одно и то же значение
	var wg sync.WaitGroup
	var wrong atomic.Int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if v, ok := m.Take(100 * time.Millisecond); !ok || v != "payload" {
				wrong.Add(1)
			}
		}()
	}
	wg.Wait()
	check(wrong.Load() == 0, "повторные Take должны видеть опубликованное значение")

	empty := newMailbox()
	_, ok = empty.Take(10 * time.Millisecond)
	check(!ok, "Take без Put должен вернуть false по таймауту")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
