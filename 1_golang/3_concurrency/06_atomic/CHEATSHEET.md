# Пакет sync/atomic — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии Go указаны там, где поведение зависит от них.

## Типы и методы

Подробно: [01_basics.md](./01_theory/01_basics.md), [02_add_load_store.md](./01_theory/02_add_load_store.md).

| Тип (Go 1.19+) | Для чего |
| :--- | :--- |
| `atomic.Int32`, `atomic.Int64` | счётчики, в том числе с вычитанием |
| `atomic.Uint32`, `atomic.Uint64` | счётчики, версии/поколения (`gen.Add(1)`) |
| `atomic.Bool` | флаги: ready, shuttingDown |
| `atomic.Pointer[T]` | снимок неизменяемого состояния (copy-on-write) |
| `atomic.Value` | то же для `any` (до дженериков) |

| Метод | Возвращает |
| :--- | :--- |
| `n.Add(d)` | **новое** значение; `d` может быть отрицательным (для знаковых типов) |
| `n.Load()` | текущее значение |
| `n.Store(v)` | — |
| `n.Swap(v)` | **старое** значение |
| `n.CompareAndSwap(old, new)` | `true`, если было `old` и заменили |

| Старый процедурный стиль | Нюанс |
| :--- | :--- |
| `atomic.AddInt64(&n, 1)`, `LoadInt64`, `StoreInt64`, `CompareAndSwapInt64` | на 32-бит платформах `int64` должен быть выровнен по 8 байт (первое поле структуры); типы `atomic.Int64` выравнивают сами |
| `float64` атомарно | только через биты: `atomic.StoreUint64(&bits, math.Float64bits(v))` |

## Порядок стоимости

Подробно: [02_add_load_store.md](./01_theory/02_add_load_store.md), [03_compare_swap.md](./01_theory/03_compare_swap.md).

| Операция | Примерно |
| :--- | :--- |
| `Load` / `Store` | ~1 ns |
| `Add` / `Swap` | ~5–15 ns |
| CAS | дороже `Add`, дешевле mutex; при высокой конкуренции CAS-loop крутится |
| `Mutex` Lock+Unlock без конкуренции | ~25–50 ns, больше при contention |

Цифры зависят от CPU и нагрузки — мерить бенчмарком.

## atomic или нет

Подробно: [04_when_not_use.md](./01_theory/04_when_not_use.md), [06_use_cases.md](./01_theory/06_use_cases.md).

| Задача | Выбор |
| :--- | :--- |
| Счётчик метрик, флаг, «последнее значение» | atomic |
| Конфиг/кэш обновляется целиком, читают часто | `atomic.Pointer[T]` + copy-on-write |
| Ленивая инициализация без mutex | `atomic.Pointer` + `CompareAndSwap(nil, s)` (или `sync.Once`) |
| «Увеличь, если меньше лимита» | CAS-loop (редко) или mutex (чаще читаемее) |
| Несколько полей должны меняться согласованно | mutex или один `atomic.Pointer` на всю структуру |
| Обновление + побочные эффекты (лог, метрика) | mutex |
| map, slice, struct целиком | mutex / `RWMutex` / `sync.Map` / `atomic.Pointer` на копию |
| Ждать события / результата | каналы |

```go
package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrLimit = errors.New("limit")

// увеличить, если меньше limit — CAS-loop
func incUpTo(n *atomic.Int64, limit int64) error {
	for {
		v := n.Load()
		if v >= limit {
			return ErrLimit
		}
		if n.CompareAndSwap(v, v+1) {
			return nil
		}
	}
}

type Config struct{ Version int }

func main() {
	var n atomic.Int64 // Go 1.19+, выравнивание автоматически
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() { defer wg.Done(); _ = incUpTo(&n, 10) }()
	}
	wg.Wait()
	fmt.Println(n.Load()) // 10

	var cfg atomic.Pointer[Config] // copy-on-write snapshot
	cfg.Store(&Config{Version: 1})
	next := *cfg.Load() // копия
	next.Version++
	cfg.Store(&next)
	fmt.Println(cfg.Load().Version) // 2

	var ready atomic.Bool
	ready.Store(true)
	fmt.Println(ready.Load(), n.Add(-1), n.Swap(0), n.Load()) // true 9 9 0
}
```

## Модель памяти

Подробно: [05_memory_model.md](./01_theory/05_memory_model.md), [go.dev/ref/mem](https://go.dev/ref/mem).

| Happens-before устанавливает | |
| :--- | :--- |
| `go f()` | → первая инструкция `f` |
| отправка в канал | → соответствующий приём |
| `close(ch)` | → приём, вернувший «закрыт» |
| `mu.Unlock()` | → следующий `mu.Lock()` того же mutex |
| atomic-запись | → atomic-чтение, увидевшее это значение (atomic-операции последовательно согласованы, формально с Go 1.19) |
| нет ни одного из них | нет гарантий порядка и видимости → data race |

## Ловушки

- ⚠️ `counter.Store(counter.Load() + 1)` — не атомарно; нужен `Add` или CAS.
- ⚠️ Смешивание atomic и обычного доступа (`n++`, `v := n`) к одной переменной — data race.
- ⚠️ Два `Load` подряд могут вернуть разные значения.
- ⚠️ Два отдельных `Add` к `count` и `total` — другая горутина видит полусостояние.
- ⚠️ `Add(-1)` на `Uint32`/`Uint64` так не работает — для вычитания берите знаковый тип.
- ⚠️ `var done bool` + `for !done {}` может крутиться вечно: видимость не гарантирована → `atomic.Bool` или канал.
- ⚠️ ABA: CAS не видит смены A→B→A (важно при переиспользовании объектов).
- ⚠️ Дорогое вычисление в CAS-loop при высокой конкуренции — ограничить попытки или взять mutex.
- ⚠️ «Lock-free ради скорости» без профиля — mutex часто не медленнее и проще.
