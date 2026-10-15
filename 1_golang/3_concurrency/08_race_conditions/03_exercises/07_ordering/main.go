// Упражнение 07: happens-before через закрытие канала.
//
// Задача: Mailbox передаёт строку от одной горутины к другой.
//   - Put(s) записывает data и ПОСЛЕ этого закрывает ready;
//   - Take(timeout) ждёт закрытия ready (не дольше timeout) и возвращает
//     data, либо ("", false).
//
// Почему это корректно: по модели памяти Go закрытие канала
// «happens before» чтения, которое увидело закрытие. Поэтому запись data
// до close гарантированно видна после <-ready. Проверка через обычный
// bool-флаг без синхронизации такой гарантии НЕ даёт — это гонка.
//
// Запуск: go run ./03_exercises/07_ordering/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Mailbox — одноразовая передача значения между горутинами.
type Mailbox struct {
	data  string
	ready chan struct{}
}

func newMailbox() *Mailbox { return &Mailbox{ready: make(chan struct{})} }

// Put публикует значение (вызывается один раз).
func (m *Mailbox) Put(s string) {
	// TODO: m.data = s; close(m.ready)
}

// Take ждёт публикации не дольше timeout.
func (m *Mailbox) Take(timeout time.Duration) (string, bool) {
	// TODO: select по m.ready и time.After(timeout).
	return "", false
}

func main() {
	fmt.Println("--- Проверка 07_ordering ---")

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
