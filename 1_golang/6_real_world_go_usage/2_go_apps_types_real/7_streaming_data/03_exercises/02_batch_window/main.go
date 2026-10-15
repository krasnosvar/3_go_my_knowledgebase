// Упражнение 02: батчинг по размеру и по времени.
//
// Писать в БД по одному сообщению дорого — копим батч и сбрасываем,
// когда набрался maxSize ИЛИ с первого сообщения в батче прошло maxWait.
//
// Задача: реализуйте (b *batcher) Add(msg string, now int64) и (b *batcher) Tick(now int64):
//   - Add: добавить сообщение (запомнить время первого в батче);
//     если len == maxSize → flush;
//   - Tick: если батч не пуст и now - firstAt >= maxWait → flush;
//   - flush: передать КОПИЮ батча в b.flushed и очистить батч.
//
// Время — тики (int64), чтобы тест был детерминированным.
//
// Запуск: go run ./03_exercises/02_batch_window
package main

import "fmt"

type batcher struct {
	maxSize int
	maxWait int64
	batch   []string
	firstAt int64
	flushed [][]string
}

func (b *batcher) flush() {
	// TODO: скопируйте b.batch в b.flushed и очистите.
}

func (b *batcher) Add(msg string, now int64) {
	// TODO: запомните время первого, добавьте, при maxSize — flush.
}

func (b *batcher) Tick(now int64) {
	// TODO: flush по времени.
}

func main() {
	fmt.Println("--- Проверка 02_batch_window ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	b := &batcher{maxSize: 3, maxWait: 10}
	b.Add("a", 0)
	b.Add("b", 1)
	b.Add("c", 2) // полный батч
	b.Add("d", 5)
	b.Tick(14) // 14-5 = 9 < 10 — рано
	check(fmt.Sprint(b.flushed) == "[[a b c]]", "сброс по размеру: "+fmt.Sprint(b.flushed))
	b.Tick(15) // прошло 10 тиков с первого в батче
	check(fmt.Sprint(b.flushed) == "[[a b c] [d]]", "сброс по времени: "+fmt.Sprint(b.flushed))
	b.Tick(100)
	check(len(b.flushed) == 2, "пустой батч не сбрасывается")
	b.Add("e", 101)
	check(len(b.flushed) > 0 && fmt.Sprint(b.flushed[0]) == "[a b c]", "сброшенные батчи — копии, не перетираются")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
