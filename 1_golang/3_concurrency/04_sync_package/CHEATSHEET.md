# Пакет sync — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии Go указаны там, где поведение зависит от них.

## Что выбрать

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

| Задача | Примитив |
| :--- | :--- |
| Простой счётчик / флаг | `sync/atomic` (см. [06_atomic](../06_atomic/CHEATSHEET.md)) |
| Защитить несколько полей вместе | `sync.Mutex` |
| Чтений намного больше записей (≈ 10:1) и секция заметна по времени | `sync.RWMutex` — после бенчмарка |
| Дождаться группы горутин | `sync.WaitGroup` (с результатом/ошибкой — `errgroup`) |
| Выполнить один раз (ленивая инициализация) | `sync.Once`, `sync.OnceValue` (Go 1.21+) |
| Один раз на ключ | `map[K]*sync.Once` или `singleflight.Group` |
| Ждать сложного условия | `sync.Cond` (редко; обычно хватает каналов и `context`) |
| Переиспользовать временные буферы в hot path | `sync.Pool` |
| Concurrent map: много чтений и мало вставок / каждая горутина пишет свои ключи | `sync.Map`; иначе `map` + `Mutex` проще и часто быстрее |
| Координация горутин, передача данных | каналы / `context` |

## WaitGroup

Подробно: [01_waitgroup.md](./01_theory/01_waitgroup.md).

| Метод | Смысл |
| :--- | :--- |
| `wg.Add(n)` | +n к счётчику — **до** `go` |
| `wg.Done()` | `Add(-1)` — в `defer` первой строкой горутины |
| `wg.Wait()` | блок, пока счётчик > 0 |
| `wg.Go(f)` (Go 1.25+) | `Add(1)` + `go` + `Done()` за один вызов |

## Mutex и RWMutex

Подробно: [02_mutex.md](./01_theory/02_mutex.md), [03_rwmutex.md](./01_theory/03_rwmutex.md).

| Метод | Кто держит |
| :--- | :--- |
| `mu.Lock()` / `mu.Unlock()` | один писатель (эксклюзивно) |
| `mu.RLock()` / `mu.RUnlock()` (`RWMutex`) | много читателей одновременно, писатель ждёт |
| Шаблон | `mu.Lock(); defer mu.Unlock()` сразу после захвата |
| Где хранить | полем рядом с защищаемыми данными; тип передавать как `*T` |

| `Mutex` | `RWMutex` |
| :--- | :--- |
| дешевле на короткой секции | выигрывает только при read-heavy и длинной секции |
| запись и чтение сравнимы по частоте | писатель может «голодать» при потоке читателей |

## Once

Подробно: [04_once.md](./01_theory/04_once.md).

| Задача | Код |
| :--- | :--- |
| Ленивая инициализация | `once.Do(func() { cfg = loadConfig() })` |
| Функция-обёртка «один раз» (Go 1.21+) | `init := sync.OnceFunc(f)` |
| Одно значение (Go 1.21+) | `get := sync.OnceValue(loadConfig)` |
| Значение + ошибка (Go 1.21+) | `get := sync.OnceValues(func() (T, error) {...})` |
| Повторная инициализация | нельзя: `Once` не сбрасывается — нужен новый |
| `f` запаниковал | `Once` считается выполненным, повторно `f` не вызовется |

## Cond и Pool

Подробно: [05_cond_pool.md](./01_theory/05_cond_pool.md).

| API | Смысл |
| :--- | :--- |
| `cond := sync.NewCond(&mu)` | условная переменная поверх mutex |
| `cond.Wait()` | под `mu.Lock()`, **в цикле** `for !ready { cond.Wait() }`: отпускает mu, ждёт, снова захватывает |
| `cond.Signal()` / `cond.Broadcast()` | разбудить одну / всех |
| `sync.Pool{New: func() any { return new(bytes.Buffer) }}` | пул временных объектов |
| `p.Get().(*bytes.Buffer)` → `b.Reset()` → работа → `p.Put(b)` | сбрасывать состояние объекта |
| Семантика `Pool` | GC может очистить в любой момент — не хранилище, а ускоритель аллокаций |

```go
package main

import (
	"bytes"
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex // не копировать: передавать *Counter
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

var bufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}

func main() {
	var c Counter
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	fmt.Println(c.n) // 100

	calls := 0
	load := sync.OnceValue(func() string { calls++; return "cfg" }) // Go 1.21+
	fmt.Println(load(), load(), calls)                               // cfg cfg 1

	b := bufPool.Get().(*bytes.Buffer)
	b.Reset() // сбросить состояние перед использованием
	b.WriteString("hi")
	fmt.Println(b.String()) // hi
	bufPool.Put(b)
}
```

## Ловушки

- ⚠️ Копирование `Mutex`, `WaitGroup`, `Once`, `Pool` (приёмник-значение, присваивание структуры) ломает внутреннее состояние; `go vet` (copylocks) ловит не всё.
- ⚠️ `Unlock` без `Lock` → `fatal error: sync: unlock of unlocked mutex` (не паника: `recover` не спасёт).
- ⚠️ `Mutex` не реентерабелен: повторный `Lock` в той же горутине — deadlock.
- ⚠️ Вложенный `RLock` в той же горутине может зависнуть, если между ними встал писатель.
- ⚠️ I/O и долгие вызовы под замком делают систему последовательной: скопировать данные, отпустить, потом работать.
- ⚠️ `wg.Add` внутри горутины или параллельно с `Wait` — гонка.
- ⚠️ `return` до `defer wg.Done()` → `Wait()` никогда не вернётся.
- ⚠️ `cond.Wait()` в `if`, а не в `for` — возможен ложный wakeup.
- ⚠️ `cond.Wait()` без захваченного mutex → `fatal error: sync: unlock of unlocked mutex` (не перехватывается `recover`).
- ⚠️ `Broadcast` будит всех — все бросаются за mutex.
- ⚠️ `Pool` для долгоживущих объектов или объектов с важным состоянием между `Get`/`Put` — ошибка.
- ⚠️ `Pool` без профиля — преждевременная оптимизация.
