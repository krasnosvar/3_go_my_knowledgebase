// Упражнение 02: батчинг по размеру и по времени (Решение).
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
	out := make([]string, len(b.batch))
	copy(out, b.batch)
	b.flushed = append(b.flushed, out)
	b.batch = b.batch[:0]
}

func (b *batcher) Add(msg string, now int64) {
	if len(b.batch) == 0 {
		b.firstAt = now
	}
	b.batch = append(b.batch, msg)
	if len(b.batch) >= b.maxSize {
		b.flush()
	}
}

func (b *batcher) Tick(now int64) {
	if len(b.batch) > 0 && now-b.firstAt >= b.maxWait {
		b.flush()
	}
}

func main() {
	fmt.Println("--- Проверка 02_batch_window (Решение) ---")

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
