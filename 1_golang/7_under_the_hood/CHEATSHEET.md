# Go под капотом — шпаргалка

Подробности и объяснения — в теории каждой темы (ссылки в разделах). ⚠️ — частая ловушка.
Версии указаны там, где поведение зависит от них; проверено на Go 1.26.

## 01. Runtime и планировщик

Подробно: [01_runtime_overview.md](./01_runtime/01_theory/01_runtime_overview.md),
[02_scheduler_gmp.md](./01_runtime/01_theory/02_scheduler_gmp.md),
[03_stack_and_preemption.md](./01_runtime/01_theory/03_stack_and_preemption.md),
[05_runtime_deep_dive.md](./01_runtime/01_theory/05_runtime_deep_dive.md).

| Сущность | Что это | Сколько |
| :--- | :--- | :--- |
| G | горутина: свой стек + состояние | миллионы |
| M | поток ОС, выполняет код | растёт при блокирующих syscalls |
| P | слот исполнения с локальной run queue | `GOMAXPROCS` |

| Механизм | Как работает |
| :--- | :--- |
| Запуск G | нужен M, держащий P; G берётся из локальной очереди P |
| Work stealing | P с пустой очередью ворует G у другого P |
| Блокирующий syscall (файлы) | M блокируется, его P уходит другому (новому) M |
| Сеть | netpoller (`epoll`/`kqueue`/IOCP): G паркуется, M не блокируется |
| Таймеры, `sleep`, каналы | G паркуется, M свободен |
| Вытеснение | в точках вызова функций + асинхронно сигналом `SIGURG` (1.14+) — `for {}` больше не захватывает P навсегда |
| Стек горутины | старт ~2 KB, растёт копированием до 1 GB (64-бит; `debug.SetMaxStack`) |
| Старт программы | `rt0_go` → инициализация runtime → `runtime.main` → все `init()` → `main.main()` |
| Конец программы | возврат из `main.main` завершает процесс, другие горутины не ждут |

| Задача | Код |
| :--- | :--- |
| Число горутин | `runtime.NumGoroutine()` |
| Ядра / P | `runtime.NumCPU()`; `runtime.GOMAXPROCS(0)` (0 — только прочитать) |
| Принудительный GC | `runtime.GC()` |
| Статистика памяти | `var ms runtime.MemStats; runtime.ReadMemStats(&ms)` |
| Дамп стеков всех G | `runtime.Stack(buf, true)` |
| Уступить P | `runtime.Gosched()` — почти не нужен с 1.14 |

| `GOMAXPROCS` по умолчанию | Версия |
| :--- | :--- |
| `runtime.NumCPU()` (видит все ядра хоста, игнорирует CPU-квоту cgroup) | ≤ 1.24 |
| учитывает CPU affinity и квоту cgroup, периодически обновляется | 1.25+ (если `go.mod` ≥ 1.25) |
| вернуть старое поведение | `GODEBUG=containermaxprocs=0`; явное значение — env `GOMAXPROCS=N` |

## 02. Память и GC

Подробно: [04_escape_analysis.md](./01_runtime/01_theory/04_escape_analysis.md),
[01_stack_vs_heap.md](./02_memory_gc/01_theory/01_stack_vs_heap.md),
[02_gc_cycle.md](./02_memory_gc/01_theory/02_gc_cycle.md),
[03_optimization.md](./02_memory_gc/01_theory/03_optimization.md).

| | Stack | Heap |
| :--- | :--- | :--- |
| Выделение | сдвиг указателя, почти бесплатно | через аллокатор (size classes, кеш per-P) |
| Освобождение | при выходе из функции | GC |
| Нагрузка на GC | нет | каждый объект — работа для mark |

| Причина escape в heap | Пример |
| :--- | :--- |
| Возврат указателя на локальную | `u := User{}; return &u` → `moved to heap: u` |
| Захват горутиной / долгоживущим замыканием | `go func() { fmt.Println(x) }()` |
| Сохранение в глобальную / долгоживущую структуру | `global = append(global, &x)` |
| Упаковка в interface | `fmt.Println(x)`, `var a any = x` (часто, не всегда) |
| Размер неизвестен или слишком велик | `make([]byte, n)` с динамическим `n`; огромный массив |

| Фаза GC (concurrent tri-color mark-and-sweep) | STW? |
| :--- | :--- |
| Sweep termination + mark setup (включение write barrier) | да, микросекунды |
| Concurrent mark (white → grey → black от roots) | нет |
| Mark termination | да, обычно < 1 мс |
| Concurrent sweep (освобождение white) | нет |

| Настройка | Эффект |
| :--- | :--- |
| `GOGC=100` (default) | следующий GC, когда heap вырос на 100% от живого после прошлого |
| `GOGC=200` / `GOGC=50` | реже GC и больше памяти / чаще GC и меньше памяти |
| `GOGC=off` | GC выключен (только эксперименты; с `GOMEMLIMIT` — GC лишь у лимита) |
| `GOMEMLIMIT=4GiB` (1.19+) | мягкий лимит памяти: у лимита GC становится агрессивнее; в K8s ставить ниже лимита pod'а |
| `debug.SetGCPercent(n)` / `debug.SetMemoryLimit(n)` | то же из кода |
| Green Tea GC | экспериментально в 1.25 (`GOEXPERIMENT=greenteagc`), по умолчанию с 1.26 |

| Источник аллокаций | Замена |
| :--- | :--- |
| `fmt.Sprintf("user_%d", id)` в hot path | `"user_" + strconv.Itoa(id)` |
| `append` в слайс без cap | `make([]T, 0, n)` |
| `string(b)` / `[]byte(s)` туда-обратно | `strings.Builder`, `bytes.Buffer` |
| Короткоживущие буферы | `sync.Pool` (`Get` → `Reset` → `Put`; пул может опустеть на любом GC) |
| Поля `bool, int64, bool` | сортировать поля по убыванию размера: 24 → 16 байт |

```go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Bad struct {
	a bool
	b int64
	c bool
}

type Good struct {
	b int64
	a bool
	c bool
}

type User struct {
	Name string `json:"name" db:"user_name"`
	age  int
}

func main() {
	fmt.Println(unsafe.Sizeof(Bad{}), unsafe.Sizeof(Good{})) // 24 16 (amd64/arm64)
	fmt.Println(unsafe.Offsetof(Bad{}.b))                    // 8

	u := User{Name: "ann"}
	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("Name")
	fmt.Println(t.Kind(), f.Tag.Get("db")) // struct user_name

	v := reflect.ValueOf(u)
	fmt.Println(v.Field(0).CanSet()) // false: копия, не адресуемо

	pv := reflect.ValueOf(&u).Elem()
	fmt.Println(pv.Field(1).CanSet()) // false: неэкспортируемое поле
	pv.Field(0).SetString("bob")
	fmt.Println(u.Name) // bob
}
```

## 03. Компилятор, модули, layout

Подробно: [01_go_src_file_structure.md](./03_compiler/01_go_src_file_structure.md),
[02_modules.md](./03_compiler/02_modules.md).

| Путь | Что там |
| :--- | :--- |
| `$GOROOT/bin` | `go`, `gofmt` |
| `$GOROOT/src` | stdlib, `runtime/` (`proc.go` — планировщик, `mgc.go` — GC, `malloc.go`, `chan.go`, `netpoll.go`), `cmd/` |
| `$GOROOT/pkg/tool/<GOOS>_<GOARCH>` | `compile`, `link`, `asm`, `cgo` (`go env GOTOOLDIR`) |
| `$GOROOT/lib/time/zoneinfo.zip` | база часовых поясов |
| `$GOPATH/bin` | `go install` |
| `$GOMODCACHE` (`$GOPATH/pkg/mod`) | кеш модулей, read-only |
| `$GOCACHE` (`~/.cache/go-build`) | кеш сборки |

| Задача | Команда |
| :--- | :--- |
| Окружение | `go env`, `go env GOROOT GOPATH GOCACHE GOMODCACHE` |
| Записать переменную | `go env -w GOPROXY=https://proxy.golang.org,direct` |
| Показать вызовы compile/link | `go build -x` |
| Escape analysis / инлайнинг | `go build -gcflags=-m ./...` (`-m -m` — подробнее) |
| Кросс-компиляция | `GOOS=linux GOARCH=arm64 go build` |
| Список платформ | `go tool dist list` |
| Символы / дизассемблер | `go tool nm bin`, `go tool objdump bin` |
| Зависимости пакета | `go list -f '{{.Deps}}' .` |
| Новый модуль | `go mod init github.com/user/project` |
| Синхронизировать `go.mod`/`go.sum` с импортами | `go mod tidy` |
| Добавить / обновить / версия / коммит / удалить | `go get pkg`, `go get -u ./...`, `go get pkg@v1.9.0`, `go get pkg@abc1234`, `go get pkg@none` |
| Только patch-обновления | `go get -u=patch ./...` |
| Скачать в кеш (CI) | `go mod download` |
| Vendor | `go mod vendor`; `go build -mod=vendor` |
| Все модули / версии модуля | `go list -m all`, `go list -m -versions pkg` |
| Почему нужен / граф / проверка | `go mod why pkg`, `go mod graph`, `go mod verify` |
| Локальная замена | `go mod edit -replace=github.com/user/lib=../lib` |
| Поднять версию языка | `go mod edit -go=1.22` |
| Очистка кешей | `go clean -modcache`, `go clean -cache -testcache` |
| Приватные модули | `GOPRIVATE=github.com/mycompany/*` (+ SSH `insteadOf` или `~/.netrc`) |

| Директива `go.mod` | Назначение |
| :--- | :--- |
| `module` | путь модуля (обязательно); v2+ — суффикс `/v2` в пути и импортах |
| `go` | минимальная версия языка (влияет на семантику, напр. loopvar 1.22) |
| `require` | зависимости (`// indirect` — транзитивные) |
| `replace` | подмена модуля / версии (локальная разработка) |
| `exclude` / `retract` | запрет версии / отзыв своей версии |
| `tool` (1.24+) | зависимости-инструменты, запуск `go tool <name>` |

Псевдоверсия коммита без тега: `v0.0.0-20230515120000-abc123def456` (базовая версия, время UTC, хеш). `go.sum` — коммитить, руками не править. `internal/` — импорт только внутри модуля; `cmd/<app>/main.go` — бинарники.

## 04. Профилирование

Подробно: [01_pprof.md](./04_profiling/01_theory/01_pprof.md), [02_profile_types.md](./04_profiling/01_theory/02_profile_types.md).

| Симптом | Профиль | Endpoint `/debug/pprof/...` |
| :--- | :--- | :--- |
| Высокий CPU | CPU (семплирование 100 Гц) | `profile?seconds=30` |
| Растёт память | heap (`inuse_space` по умолчанию) | `heap` |
| Горячий аллокатор | все аллокации с старта (`alloc_space`) | `allocs` |
| Висит / растёт `NumGoroutine` | стеки всех горутин | `goroutine` (`?debug=1` — сгруппировано, `?debug=2` — полностью) |
| Низкий параллелизм, борьба за lock | mutex (выключен по умолчанию) | `mutex` + `runtime.SetMutexProfileFraction(1)` |
| Долгие блокировки на каналах / lock | block (выключен по умолчанию) | `block` + `runtime.SetBlockProfileRate(1)` |
| Странные паузы, планировщик, GC | execution trace | `trace?seconds=5` → `go tool trace trace.out` |

| Задача | Команда |
| :--- | :--- |
| Включить HTTP-профайлер | `import _ "net/http/pprof"` + HTTP-сервер (`DefaultServeMux`), не на публичном порту |
| CPU с живого сервиса | `go tool pprof http://localhost:8080/debug/pprof/profile?seconds=5` |
| Web UI (нужен Graphviz для графа) | `go tool pprof -http=:8081 <url или файл>` |
| Профили из бенчмарка | `go test -bench=. -benchmem -cpuprofile cpu.out -memprofile mem.out` |
| Trace из теста | `go test -trace trace.out` → `go tool trace trace.out` |
| Без интерактива | `go tool pprof -top cpu.out` |
| Режим heap | `-sample_index=inuse_space` / `inuse_objects` / `alloc_space` / `alloc_objects` |

| Команда в `(pprof)` | Что показывает |
| :--- | :--- |
| `top` / `top10` | самые тяжёлые функции |
| `list <func>` | исходник с затратами по строкам |
| `traces` | цепочки вызовов |
| `web` | граф в браузере |

`flat` — затраты самой функции; `cum` — функции вместе со всем, что она вызвала. Утечка: растут и `inuse_space`, и `alloc_space`; горячий аллокатор: растёт только `alloc_space`.

| `GODEBUG` | Вывод в stderr |
| :--- | :--- |
| `gctrace=1` | строка на каждый GC: `gc 1 @0.004s 2%: 0.075+0.48+0.084 ms clock, ... 3->4->1 MB, 4 MB goal, ... 22 P` |
| `schedtrace=1000` (+ `scheddetail=1`) | состояние P/M/G раз в секунду |
| `inittrace=1` | время и аллокации каждого `init` |

## 05. Reflection и unsafe

Подробно: [01_reflection.md](./05_reflection_unsafe/01_reflection.md), [02_unsafe.md](./05_reflection_unsafe/02_unsafe.md).

| Закон рефлексии | Код |
| :--- | :--- |
| 1. interface → объект рефлексии | `t := reflect.TypeOf(x)`; `v := reflect.ValueOf(x)` |
| 2. объект рефлексии → interface | `v.Interface().(T)` |
| 3. менять можно только settable | `v := reflect.ValueOf(&x).Elem()`; `v.CanSet()`; `v.SetInt(5)` |

| Задача | Код |
| :--- | :--- |
| Вид типа | `t.Kind() == reflect.Struct` (тип `User` → Kind `Struct`) |
| Поля структуры | `t.NumField()`, `t.Field(i)`, `t.FieldByName("Name")`, `v.Field(i)` |
| Тег | `f.Tag.Get("json")`, `f.Tag.Lookup("db")` |
| Методы | `v.MethodByName("Do").Call([]reflect.Value{reflect.ValueOf(arg)})` |
| Сигнатура функции | `t.NumIn()`, `t.In(i)`, `t.NumOut()`, `t.IsVariadic()` |
| Создать | `reflect.New(t)` (указатель), `reflect.Zero(t)`, `reflect.MakeSlice`, `reflect.MakeMap`, `reflect.MakeFunc` |
| Map | `v.SetMapIndex(k, val)`, `v.MapIndex(k)`, `v.MapRange()` |
| Проверки | `IsValid()`, `CanSet()`, `CanAddr()`, `CanInterface()`, `IsNil()` |
| Разыменовать | `v.Elem()`, `reflect.Indirect(v)` |
| Интерфейс реализован? | `t.Implements(reflect.TypeOf((*error)(nil)).Elem())` |
| Сравнить что угодно | `reflect.DeepEqual(a, b)` |

| `unsafe` | Назначение |
| :--- | :--- |
| `unsafe.Sizeof` / `Alignof` / `Offsetof` | размер, выравнивание, смещение поля |
| `unsafe.Pointer` | конверсия между типами указателей, арифметика через `uintptr` в одном выражении |
| `unsafe.Slice(ptr, n)` (1.17+) / `unsafe.SliceData(s)` (1.20+) | слайс из указателя / указатель на массив слайса |
| `unsafe.String(ptr, n)` / `unsafe.StringData(s)` (1.20+) | строка без копии / указатель на байты строки |
| Альтернативы | `encoding/binary`, `atomic.Pointer[T]`, generics, cgo |

## Ловушки

- ⚠️ Бесконечная рекурсия — не recoverable panic, а `fatal error: stack overflow` после роста стека до 1 GB.
- ⚠️ Файловый I/O блокирует M: тысячи горутин на диске → тысячи потоков ОС.
- ⚠️ Тысячи горутин на одном mutex — contention, а не параллелизм.
- ⚠️ До 1.25 в контейнере `GOMAXPROCS` = все ядра хоста → throttling по CPU-квоте.
- ⚠️ `GOMEMLIMIT` из `/proc/meminfo` в pod'е — это память узла, а не лимит контейнера (`/sys/fs/cgroup/memory.max`).
- ⚠️ `GOMEMLIMIT` ниже живого heap → GC работает почти непрерывно (death spiral по CPU).
- ⚠️ `GODEBUG=allocfreetrace` удалён из runtime — в Go 1.26 не работает.
- ⚠️ Оптимизация памяти без `pprof` / `-benchmem` — время впустую.
- ⚠️ `sync.Pool` — не кеш: содержимое пропадает на GC; объект перед `Put` надо сбрасывать.
- ⚠️ mutex/block-профили пустые, пока не включены `SetMutexProfileFraction` / `SetBlockProfileRate`.
- ⚠️ `net/http/pprof` на публичном порту — утечка внутренностей и DoS-вектор.
- ⚠️ v2+ модуля без `/v2` в пути `module` и импортах — не работает.
- ⚠️ `go get -u ./...` обновляет всё дерево разом — проверяйте тесты.
- ⚠️ `reflect.ValueOf(x).SetInt(...)` без указателя → panic `using unaddressable value`.
- ⚠️ Неэкспортируемые поля через reflect читать можно, менять нельзя (`CanSet() == false`).
- ⚠️ Рефлексия в 10–100 раз медленнее обычного кода и переносит ошибки типов в runtime.
- ⚠️ `unsafe.String(&b[0], len(b))` — изменение `b` после этого ломает неизменяемость строки; пустой `b` → panic.
- ⚠️ `uintptr`, сохранённый в переменную, не держит объект для GC → dangling pointer.
