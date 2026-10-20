// Пример 05: доставка at-least-once + идемпотентный консьюмер.
//
// Продюсер имитирует ненадёжную доставку: некоторые сообщения он шлёт
// ПОВТОРНО (дубликаты) — как будто ack потерялся и брокер прислал заново.
// Консьюмер идемпотентен: он дедуплицирует по message-id (карта seen) и
// применяет бизнес-эффект (складывает суммы в balance) ровно один раз на id.
//
// Итог: несмотря на дубли в потоке доставки, эффект применён однократно,
// и итоговый баланс совпадает с суммой уникальных сообщений.
//
// Демонстрирует: at-least-once (дубли), дедупликацию по id, идемпотентность.
package main

import (
	"fmt"
	"sync"
)

// message — сообщение с уникальным id (ключ идемпотентности) и полезной
// нагрузкой.
type message struct {
	id     string
	amount int
}

func main() {
	fmt.Println("--- 02_at_least_once_idempotent ---")

	// Уникальные сообщения (то, что реально произошло в бизнесе).
	unique := []message{
		{"m1", 10},
		{"m2", 20},
		{"m3", 5},
	}

	// Поток доставки at-least-once: m1 и m3 доставлены дважды (дубли).
	delivery := []message{
		{"m1", 10},
		{"m2", 20},
		{"m1", 10}, // дубль
		{"m3", 5},
		{"m3", 5}, // дубль
	}

	ch := make(chan message, len(delivery))

	var wg sync.WaitGroup

	// Идемпотентный консьюмер.
	seen := make(map[string]bool) // множество уже обработанных id
	balance := 0                  // бизнес-эффект
	applied := 0                  // сколько раз реально применили
	skipped := 0                  // сколько дублей отсеяли

	wg.Add(1)
	go func() {
		defer wg.Done()
		for m := range ch {
			if seen[m.id] {
				skipped++
				fmt.Printf("дубль отброшен: id=%s\n", m.id)
				continue
			}
			seen[m.id] = true
			balance += m.amount
			applied++
			fmt.Printf("применено: id=%s amount=%d (balance=%d)\n", m.id, m.amount, balance)
		}
	}()

	// «Брокер» доставляет сообщения (с дублями) и закрывает канал.
	for _, m := range delivery {
		ch <- m
	}
	close(ch)

	wg.Wait()

	// Ожидаемый баланс — сумма уникальных сообщений.
	want := 0
	for _, m := range unique {
		want += m.amount
	}

	fmt.Printf("доставлено=%d применено=%d отброшено_дублей=%d balance=%d (ожидалось %d)\n",
		len(delivery), applied, skipped, balance, want)

	if balance == want && applied == len(unique) {
		fmt.Println("PASS: эффект применён ровно один раз на каждый уникальный id")
	} else {
		fmt.Println("FAIL: эффект применён неверно")
	}
}
