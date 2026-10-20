// Пример 03: кворумное чтение/запись (W + R > N).
//
// Модель: N in-memory "реплик", у каждой хранится версионированное
// значение (version, value). Запись идёт на W реплик, чтение — с R реплик,
// из прочитанных выбирается версия с наибольшим номером.
//
// Демонстрируем ключевой факт: при W + R > N множества узлов записи и
// чтения обязаны пересекаться, поэтому чтение видит последнюю запись.
// При W + R <= N пересечение не гарантировано — чтение может вернуть
// устаревшее значение.
//
// Реплики обновляются конкурентно в горутинах; синхронизация — через
// sync.WaitGroup и sync.Mutex (никаких time.Sleep). Порядок реплик
// фиксирован, поэтому результат детерминирован.
//
// Запуск: go run ./02_examples/03_quorum_rw/main.go
package main

import (
	"fmt"
	"sort"
	"sync"
)

// replica — одна реплика с версионированным значением.
type replica struct {
	mu      sync.Mutex
	version int
	value   string
}

func (r *replica) write(version int, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Применяем только более новую версию (last-write-wins по версии).
	if version > r.version {
		r.version = version
		r.value = value
	}
}

type versioned struct {
	version int
	value   string
}

func (r *replica) read() versioned {
	r.mu.Lock()
	defer r.mu.Unlock()
	return versioned{r.version, r.value}
}

// cluster — набор из N реплик.
type cluster struct {
	replicas []*replica
}

func newCluster(n int) *cluster {
	c := &cluster{replicas: make([]*replica, n)}
	for i := range c.replicas {
		c.replicas[i] = &replica{}
	}
	return c
}

// writeQuorum пишет значение на первые w реплик (индексы 0..w-1),
// конкурентно, дожидаясь всех через WaitGroup.
func (c *cluster) writeQuorum(indices []int, version int, value string) {
	var wg sync.WaitGroup
	for _, idx := range indices {
		wg.Add(1)
		go func(rep *replica) {
			defer wg.Done()
			rep.write(version, value)
		}(c.replicas[idx])
	}
	wg.Wait()
}

// readQuorum читает с указанных реплик и возвращает значение с максимальной
// версией. Чтения конкурентны, результаты собираются под мьютексом.
func (c *cluster) readQuorum(indices []int) versioned {
	var wg sync.WaitGroup
	var mu sync.Mutex
	best := versioned{version: -1}
	for _, idx := range indices {
		wg.Add(1)
		go func(rep *replica) {
			defer wg.Done()
			got := rep.read()
			mu.Lock()
			if got.version > best.version {
				best = got
			}
			mu.Unlock()
		}(c.replicas[idx])
	}
	wg.Wait()
	return best
}

// firstN возвращает срез индексов [0, n).
func firstN(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = i
	}
	return out
}

// lastN возвращает срез индексов [total-n, total) — крайние реплики.
func lastN(total, n int) []int {
	out := make([]int, 0, n)
	for i := total - n; i < total; i++ {
		out = append(out, i)
	}
	return out
}

func intersects(a, b []int) []int {
	set := map[int]bool{}
	for _, x := range a {
		set[x] = true
	}
	var common []int
	for _, x := range b {
		if set[x] {
			common = append(common, x)
		}
	}
	sort.Ints(common)
	return common
}

// runScenario моделирует: начальная запись v1 на ВСЕ реплики, затем новая
// запись v2 на W реплик, затем чтение с R реплик. Возвращает, увидело ли
// чтение свежую v2.
func runScenario(n, w, r int) (seenValue string, seenLatest bool, writeSet, readSet, overlap []int) {
	c := newCluster(n)

	// Изначально на всех репликах версия 1 (согласованное старое состояние).
	c.writeQuorum(firstN(n), 1, "v1-old")

	// Новая запись версии 2 попадает на первые W реплик.
	writeSet = firstN(w)
	c.writeQuorum(writeSet, 2, "v2-new")

	// Чтение с последних R реплик (worst case для пересечения).
	readSet = lastN(n, r)
	res := c.readQuorum(readSet)

	overlap = intersects(writeSet, readSet)
	return res.value, res.version == 2, writeSet, readSet, overlap
}

func main() {
	fmt.Println("--- Пример: кворум W + R > N ---")

	const n = 3

	scenarios := []struct {
		w, r int
	}{
		{2, 2}, // W+R=4 > 3  -> кворум соблюдён
		{1, 1}, // W+R=2 <= 3 -> кворум НЕ соблюдён
		{3, 1}, // W+R=4 > 3  -> кворум соблюдён
	}

	allExpected := true
	for _, s := range scenarios {
		val, latest, ws, rs, ov := runScenario(n, s.w, s.r)
		quorumOK := s.w+s.r > n

		fmt.Printf("\nN=%d W=%d R=%d  (W+R=%d %s N)\n", n, s.w, s.r, s.w+s.r,
			map[bool]string{true: ">", false: "<="}[quorumOK])
		fmt.Printf("  запись v2 на реплики %v, чтение с реплик %v\n", ws, rs)
		fmt.Printf("  пересечение множеств: %v\n", ov)
		fmt.Printf("  чтение вернуло: %q (свежее=%v)\n", val, latest)

		// Проверяем инвариант: свежесть чтения совпадает с наличием кворума.
		if latest != quorumOK {
			allExpected = false
		}
	}

	fmt.Println("\nВывод: когда W+R>N, множества записи и чтения пересекаются,")
	fmt.Println("и чтение гарантированно видит последнюю запись. Когда W+R<=N,")
	fmt.Println("пересечения может не быть — читаем устаревшее значение.")

	if allExpected {
		fmt.Println("Инвариант кворума подтверждён во всех сценариях.")
	} else {
		fmt.Println("ВНИМАНИЕ: инвариант нарушен (неожиданно).")
	}
	fmt.Println("Готово, программа завершается.")
}
