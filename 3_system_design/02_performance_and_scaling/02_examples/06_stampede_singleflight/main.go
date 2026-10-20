// Пример 06: защита от cache stampede (single-flight).
//
// N параллельных горутин одновременно запрашивают ОДИН И ТОТ ЖЕ ключ,
// которого ещё нет в кэше. Без защиты все N разом пошли бы грузить его
// из «БД» (thundering herd). Здесь single-flight гарантирует, что
// загрузка выполнится РОВНО ОДИН РАЗ, а остальные горутины дождутся её
// результата.
//
// Реализация: sync.Mutex + карта in-flight вызовов. Первый по ключу
// создаёт «call» и грузит; остальные находят существующий call и ждут
// его завершения через call.wg (sync.WaitGroup), не запуская загрузку.
//
// Детерминизм: синхронизация построена на sync.WaitGroup, а не на sleep.
// Счётчик загрузок (atomic) должен быть равен 1.
//
// Запуск: go run ./02_examples/06_stampede_singleflight/main.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// call — одна выполняющаяся (или уже выполненная) загрузка ключа.
type call struct {
	wg  sync.WaitGroup
	val string
}

// group реализует single-flight: подавляет дублирующие вызовы по ключу.
type group struct {
	mu sync.Mutex
	m  map[string]*call
}

func newGroup() *group { return &group{m: make(map[string]*call)} }

// Do выполняет fn для key так, что при нескольких одновременных вызовах
// с одним key реально запустится только один fn, а остальные получат его
// результат. Возвращает результат и признак shared (значение получено от
// чужого, уже идущего вызова).
func (g *group) Do(key string, fn func() string) (val string, shared bool) {
	g.mu.Lock()
	if c, ok := g.m[key]; ok {
		// Загрузка по этому ключу уже идёт — ждём её, не запуская свою.
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, true
	}

	// Мы первые по этому ключу: регистрируем call и берём загрузку на себя.
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val = fn() // единственный реальный вызов «загрузки»
	c.wg.Done()  // будим всех ожидающих

	// Убираем ключ, чтобы следующие обращения (после того как кэш опустеет)
	// снова могли инициировать загрузку.
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, false
}

func main() {
	fmt.Println("--- Пример single-flight против cache stampede ---")

	const workers = 50
	var loads int64 // сколько раз реально выполнилась «загрузка из БД»

	g := newGroup()

	// «Медленная» загрузка: считаем вызовы и держим паузу, чтобы все
	// горутины гарантированно застали загрузку в процессе.
	load := func() string {
		atomic.AddInt64(&loads, 1)
		time.Sleep(20 * time.Millisecond)
		return "value-of-K"
	}

	// Стартовый барьер: отпускаем все горутины одновременно (детерминизм —
	// на WaitGroup, а не на угадывании таймингов).
	var start sync.WaitGroup
	start.Add(1)

	var done sync.WaitGroup
	results := make([]string, workers)
	sharedCount := int64(0)

	for i := 0; i < workers; i++ {
		done.Add(1)
		go func(idx int) {
			defer done.Done()
			start.Wait() // ждём общего старта
			val, shared := g.Do("K", load)
			results[idx] = val
			if shared {
				atomic.AddInt64(&sharedCount, 1)
			}
		}(i)
	}

	start.Done() // одновременный старт всех горутин
	done.Wait()  // дожидаемся всех

	n := atomic.LoadInt64(&loads)
	fmt.Printf("горутин: %d\n", workers)
	fmt.Printf("реальных загрузок из «БД»: %d (ожидалось 1)\n", n)
	fmt.Printf("получили результат от чужого вызова (shared): %d\n", sharedCount)

	allSame := true
	for _, r := range results {
		if r != "value-of-K" {
			allSame = false
			break
		}
	}

	if n == 1 && allSame {
		fmt.Println("Все горутины получили одно значение при единственной загрузке — stampede подавлен.")
	} else {
		fmt.Println("Что-то не так: загрузка выполнилась не один раз.")
	}
	fmt.Println("Готово, программа завершается.")
}
