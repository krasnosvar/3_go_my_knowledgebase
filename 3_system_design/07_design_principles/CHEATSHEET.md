# Принципы проектирования — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## SOLID в Go

Подробно: [01_solid.md](./01_theory/01_solid.md).

| Принцип | Суть | Как в Go | Нарушение (запах) |
| :--- | :--- | :--- | :--- |
| S — Single Responsibility | одна причина для изменения | маленькие типы: `Payroll`, `Report`, `Store` | один тип делает валидацию + БД + HTTP |
| O — Open/Closed | расширять без правки | интерфейс + новые реализации | `switch` по типу, дополняемый на каждую сущность |
| L — Liskov Substitution | любая реализация соблюдает контракт | поведение при ошибках и `nil` — часть контракта | одна реализация `Get` возвращает ошибку, другая паникует |
| I — Interface Segregation | клиент не зависит от лишних методов | `io.Reader`, `io.Writer`, композиция `ReadWriter` | «толстый» интерфейс на 15 методов, нужны два |
| D — Dependency Inversion | оба уровня зависят от абстракции | интерфейс через конструктор: `NewReport(n Notifier)` | `*sql.DB` прямо в бизнес-логике |

Цепочка: SRP → маленькие типы → узкие интерфейсы (ISP) → зависимость от абстракции (DIP) → расширение реализациями и декораторами (OCP).
Интерфейс объявляют **у потребителя**: «accept interfaces, return structs».

## DRY / KISS / YAGNI

Подробно: [02_dry_kiss_yagni.md](./01_theory/02_dry_kiss_yagni.md).

| Принцип | Про что | Пример нарушения |
| :--- | :--- | :--- |
| DRY | не дублировать **знание** (бизнес-правило) | правило «скидка 10% при сумме > 1000» в трёх местах |
| KISS | **как** писать: проще | рефлексия/кодогенерация ради пары строк, функции с десятком флагов |
| YAGNI | **что** писать: только нужное сейчас | конфиг под 5 «будущих» БД, плагины и фабрики «на вырост» |

- Ложный DRY: случайно похожий код, меняющийся по разным причинам, не объединять. «A little copying is better than a little dependency».
- DRY против KISS: убирать дублирование знания, а не гнаться за нулём повторов.

## Паттерны в идиоматичном Go

Подробно: [03_design_patterns.md](./01_theory/03_design_patterns.md).

| Группа | Паттерн | В Go |
| :--- | :--- | :--- |
| Порождающие | Factory | `func NewX(...) Iface` |
| | Builder | functional options: `NewServer(WithTimeout(5), WithTLS())` |
| | Singleton | пакетная переменная + `sync.Once` |
| | Abstract Factory / Prototype | редко, обычно избыточны |
| Структурные | Adapter | тип-обёртка, `http.HandlerFunc` (func → interface) |
| | Decorator | middleware над `http.Handler`, обёртка над интерфейсом |
| | Facade | пакет с удобным API |
| | Proxy | обёртка с той же сигнатурой (кэш, ленивость, доступ) |
| | Composite | дерево через общий интерфейс |
| Поведенческие | Strategy | func-тип или интерфейс с одним методом; `sort.Slice(s, less)` |
| | Observer | каналы (`chan`) |
| | Iterator | `range`; `iter.Seq` (Go 1.23+) |
| | Command | `func()` как значение |
| | Template Method | композиция / передача функций вместо наследования |
| | Chain of Responsibility | цепочка middleware |

```go
package main

import (
	"fmt"
	"sort"
)

type Server struct{ timeout int }
type Option func(*Server)

func WithTimeout(t int) Option { return func(s *Server) { s.timeout = t } }

func NewServer(opts ...Option) *Server {
	s := &Server{timeout: 30} // значения по умолчанию
	for _, o := range opts {
		o(s)
	}
	return s
}

type SortStrategy func(a, b int) bool // Strategy = func-тип

func main() {
	fmt.Println(NewServer().timeout, NewServer(WithTimeout(5)).timeout) // 30 5
	var desc SortStrategy = func(a, b int) bool { return a > b }
	s := []int{3, 1, 2}
	sort.Slice(s, func(i, j int) bool { return desc(s[i], s[j]) })
	fmt.Println(s) // [3 2 1]
}
```

## 12-factor app

Подробно: [04_twelve_factor.md](./01_theory/04_twelve_factor.md).

| # | Фактор | Суть | В Go |
| :--- | :--- | :--- | :--- |
| 1 | Codebase | один репозиторий → много деплоев | — |
| 2 | Dependencies | явно объявлены и изолированы | `go.mod` / `go.sum` |
| 3 | Config | в переменных окружения | `os.Getenv`, секреты не коммитить |
| 4 | Backing services | БД/кэш/очередь — ресурсы по URL из конфига | DSN из env |
| 5 | Build, release, run | стадии разделены, релиз неизменяем | один бинарник + env |
| 6 | Processes | stateless, состояние — в backing services | scratch/distroless-образ |
| 7 | Port binding | сервис сам слушает порт | `http.ListenAndServe(":"+port, mux)` |
| 8 | Concurrency | масштаб — числом процессов | реплики / HPA |
| 9 | Disposability | быстрый старт, graceful shutdown по SIGTERM | `signal.NotifyContext` + `Server.Shutdown(ctx)` |
| 10 | Dev/prod parity | окружения максимально похожи | те же backing services и версии |
| 11 | Logs | поток событий в `stdout` | `log`, `slog`; без ротации файлов в приложении |
| 12 | Admin processes | разовые задачи — одноразовые процессы того же релиза | миграции отдельной командой |

## На code review и в проде

Подробно: [05_use_cases.md](./01_theory/05_use_cases.md).

| Видите | Применить |
| :--- | :--- |
| Тест требует реальную БД/HTTP | DIP: зависимость интерфейсом, в тесте фейк |
| Сквозная функциональность (логи, auth, метрики, recover) | Decorator: `mux := WithLogging(WithAuth(handler))` |
| Сменные тарифы/форматы/политики ретраев | Strategy: func-значение или маленький интерфейс |
| Конструктор с 5+ параметрами по умолчанию | functional options |
| Один образ на dev/stage/prod | конфиг из env, логи в stdout, graceful shutdown, stateless |

## Ловушки

- ⚠️ Переносить ООП-паттерны буквально (иерархии классов-стратегий, фабрики фабрик) — сначала спросите, не хватит ли функции, интерфейса или композиции.
- ⚠️ Ложный DRY: насильно объединённые случайно похожие фрагменты создают неправильную связанность.
- ⚠️ Абстракции и плагины «на будущее» без требования в бэклоге — нарушение YAGNI.
- ⚠️ Разные реализации интерфейса по-разному ведут себя при ошибке/`nil` — нарушение LSP.
- ⚠️ Конфиг и секреты в коде или файлах репозитория вместо env — один билд нельзя катить в разные окружения.
- ⚠️ Состояние (сессии, кэш) в памяти процесса — ломает масштабирование и перезапуск (факторы 6, 8, 9).
- ⚠️ Игнор SIGTERM → оборванные запросы на каждом rolling-обновлении.
