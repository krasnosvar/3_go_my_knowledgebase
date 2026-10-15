// Упражнение 04: dead letter queue.
//
// Задача: реализуйте (w *worker) Handle(m msg):
//   - вызвать w.process(m); успех → m.ID в w.done;
//   - ошибка → m.Attempts++; если Attempts < w.maxAttempts — вернуть
//     сообщение в конец w.queue (повтор позже); иначе — в w.dlq
//     (вместе с текстом последней ошибки в поле LastErr).
//
// И (w *worker) Run(): обрабатывать w.queue, пока она не опустеет.
//
// Запуск: go run ./03_exercises/04_dlq_move
package main

import (
	"errors"
	"fmt"
)

type msg struct {
	ID       string
	Attempts int
	LastErr  string
}

type worker struct {
	queue       []msg
	dlq         []msg
	done        []string
	maxAttempts int
	process     func(msg) error
}

func (w *worker) Handle(m msg) {
	// TODO: успех → done; ошибка → повтор или DLQ.
}

func (w *worker) Run() {
	// TODO: пока очередь не пуста — снять первый и Handle.
	w.queue = nil
}

func main() {
	fmt.Println("--- Проверка 04_dlq_move ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	flaky := 0
	w := &worker{
		queue:       []msg{{ID: "ok"}, {ID: "flaky"}, {ID: "poison"}},
		maxAttempts: 3,
		process: func(m msg) error {
			switch m.ID {
			case "flaky":
				flaky++
				if flaky < 2 {
					return errors.New("timeout")
				}
			case "poison":
				return errors.New("cannot parse")
			}
			return nil
		},
	}
	w.Run()

	check(fmt.Sprint(w.done) == "[ok flaky]", "ok и flaky (со 2-й попытки) обработаны: "+fmt.Sprint(w.done))
	check(len(w.dlq) == 1 && w.dlq[0].ID == "poison", "poison ушёл в DLQ")
	if len(w.dlq) == 1 {
		check(w.dlq[0].Attempts == 3 && w.dlq[0].LastErr == "cannot parse", "в DLQ сохранены попытки и последняя ошибка")
	}
	check(len(w.queue) == 0, "очередь пуста")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
