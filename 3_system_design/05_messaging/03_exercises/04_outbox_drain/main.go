// Упражнение 04: слив transactional outbox в брокер (Хардкор).
//
// Задача: реализуйте relay, который «сливает» события из in-memory outbox в
// брокер по паттерну transactional outbox. Смоделирован сбой подтверждения:
// сообщение доходит до брокера, но ack теряется (publish возвращает false),
// поэтому событие остаётся непомеченным и будет переслано на следующем проходе
// -> получается at-least-once (в брокере появляется дубль). Дедупликация — не
// ваша забота, её делает консьюмер; здесь важно НЕ ПОТЕРЯТЬ ни одного события.
//
// Реализуйте drain(ob):
//   - идём по событиям outbox по порядку;
//   - для каждого НЕопубликованного события вызываем ob.publish(e):
//   - publish возвращает false (сбой) -> НЕ помечаем published (перешлём позже);
//   - publish возвращает true (успех)  -> помечаем e.published = true;
//   - drain делает проходы, пока не останется неопубликованных событий.
//
// В этой модели publish первого события первый раз падает (см. outbox.publish),
// поэтому drain обязан сделать повторный проход и всё же доставить его.
//
// Проверка: в брокере оказались ВСЕ события (без потерь), а первое —
// доставлено дважды (демонстрация at-least-once).
//
// Запуск: go run ./03_exercises/04_outbox_drain/main.go
package main

import "fmt"

type event struct {
	id        int
	published bool
}

// outbox — in-memory outbox + «брокер» (broker) как приёмник публикаций.
type outbox struct {
	events   []*event
	broker   []int // id событий, доставленных в брокер (могут повторяться)
	failOnce map[int]bool
}

// publish имитирует публикацию в брокер. Сообщение ВСЕГДА доходит до брокера
// (append в broker), но для событий из failOnce ПЕРВЫЙ раз теряется ack:
// publish возвращает false, будто relay не узнал об успехе. Тогда relay
// перешлёт событие ещё раз -> в брокере появится дубль. Это и есть источник
// at-least-once: сообщение доставлено, но подтверждение потеряно.
func (o *outbox) publish(e *event) bool {
	o.broker = append(o.broker, e.id) // сообщение дошло до брокера
	if o.failOnce[e.id] {
		o.failOnce[e.id] = false // ack «потерян» только в первый раз
		return false             // relay считает публикацию неуспешной -> перешлёт
	}
	return true
}

// drain сливает все неопубликованные события из outbox в брокер, повторяя
// проходы при сбоях, пока не останется неопубликованных.
func drain(o *outbox) {
	// TODO: реализуйте проходы по o.events:
	//   - пропускайте уже опубликованные;
	//   - вызывайте o.publish(e); при успехе помечайте e.published = true;
	//   - повторяйте, пока есть неопубликованные события.
}

var _ = drain

func main() {
	fmt.Println("--- Проверка 04_outbox_drain ---")

	ob := &outbox{
		events: []*event{
			{id: 1},
			{id: 2},
			{id: 3},
		},
		failOnce: map[int]bool{1: true}, // публикация события 1 падает один раз
	}

	drain(ob)

	allOK := true

	// 1) Все события помечены published.
	for _, e := range ob.events {
		if !e.published {
			allOK = false
			fmt.Printf("FAIL: событие %d осталось неопубликованным\n", e.id)
		}
	}

	// 2) В брокере есть все id (без потерь).
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

	// 3) Событие 1 доставлено дважды (at-least-once из-за повторной попытки).
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
