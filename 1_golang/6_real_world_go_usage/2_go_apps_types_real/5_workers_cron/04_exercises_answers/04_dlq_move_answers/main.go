// Упражнение 04: dead letter queue (Решение).
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
	err := w.process(m)
	if err == nil {
		w.done = append(w.done, m.ID)
		return
	}
	m.Attempts++
	m.LastErr = err.Error()
	if m.Attempts < w.maxAttempts {
		w.queue = append(w.queue, m)
		return
	}
	w.dlq = append(w.dlq, m)
}

func (w *worker) Run() {
	for len(w.queue) > 0 {
		m := w.queue[0]
		w.queue = w.queue[1:]
		w.Handle(m)
	}
}

func main() {
	fmt.Println("--- Проверка 04_dlq_move (Решение) ---")

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
