// Упражнение 08: слив transactional outbox в брокер (Решение).
package main

import "fmt"

type event struct {
	id        int
	published bool
}

type outbox struct {
	events   []*event
	broker   []int
	failOnce map[int]bool
}

func (o *outbox) publish(e *event) bool {
	o.broker = append(o.broker, e.id) // сообщение дошло до брокера
	if o.failOnce[e.id] {
		o.failOnce[e.id] = false // ack «потерян» только в первый раз
		return false
	}
	return true
}

func drain(o *outbox) {
	for {
		pending := false
		for _, e := range o.events {
			if e.published {
				continue
			}
			pending = true
			if o.publish(e) {
				e.published = true // помечаем только после успешной публикации
			}
			// при неуспехе оставляем published=false -> следующий проход
			// перешлёт
		}
		if !pending {
			return // все события опубликованы
		}
	}
}

var _ = drain

func main() {
	fmt.Println("--- Проверка 08_outbox_drain (Решение) ---")

	ob := &outbox{
		events: []*event{
			{id: 1},
			{id: 2},
			{id: 3},
		},
		failOnce: map[int]bool{1: true},
	}

	drain(ob)

	allOK := true

	for _, e := range ob.events {
		if !e.published {
			allOK = false
			fmt.Printf("FAIL: событие %d осталось неопубликованным\n", e.id)
		}
	}

	seen := map[int]int{}
	for _, id := range ob.broker {
		seen[id]++
	}
	for _, e := range ob.events {
		if seen[e.id] < 1 {
			allOK = false
			fmt.Printf("FAIL: событие %d не доставлено в брокер\n", e.id)
		}
	}

	if seen[1] != 2 {
		allOK = false
		fmt.Printf("FAIL: событие 1 доставлено %d раз, ожидалось 2 (at-least-once)\n", seen[1])
	}

	fmt.Printf("broker=%v (дубль события 1 — это ожидаемый at-least-once)\n", ob.broker)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
