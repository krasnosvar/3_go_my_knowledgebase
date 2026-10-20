# SOLID и паттерны проектирования в Go — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## SOLID в Go

Подробно: [01_solid.md](./01_theory/01_solid.md).

| Принцип | Суть | Как в Go | Нарушение (запах) |
| :--- | :--- | :--- | :--- |
| S — Single Responsibility | одна причина для изменения | маленькие типы: `Payroll`, `Report`, `Store` | один тип делает валидацию + БД + HTTP |
| O — Open/Closed | расширять без правки | интерфейс + новые реализации | `switch` по типу, дополняемый на сущность |
| L — Liskov Substitution | реализация соблюдает контракт | поведение при ошибках и `nil` — в контракте | одна реализация возвращает `error`, другая `panic` |
| I — Interface Segregation | без лишних методов у клиента | `io.Reader`, `io.Writer`, `ReadWriter` | «толстый» интерфейс на 15 методов |
| D — Dependency Inversion | оба уровня зависят от абстракции | интерфейс в конструкторе `NewReport(n Notifier)` | `*sql.DB` прямо в бизнес-логике |

Цепочка: SRP → маленькие типы → узкие интерфейсы (ISP) → зависимость от
абстракции (DIP) → расширение реализациями и декораторами (OCP).
Интерфейс объявляют **у потребителя**: «accept interfaces, return structs».

## DRY / KISS / YAGNI

Подробно: [02_dry_kiss_yagni.md](./01_theory/02_dry_kiss_yagni.md).

| Принцип | Про что | Пример нарушения |
| :--- | :--- | :--- |
| DRY | не дублировать **знание** (правило) | правило «скидка 10% при > 1000» в трёх местах |
| KISS | **как** писать: проще | рефлексия ради пары строк, 10 булевых флагов |
| YAGNI | **что** писать: только нужное сейчас | конфиг и фабрики под 5 «будущих» БД |

- Ложный DRY: случайно похожий код, меняющийся по разным причинам, не
  объединять. «A little copying is better than a little dependency».
- DRY против KISS: убирать дублирование знания, а не гнаться за нулём повторов.

## Паттерны в идиоматичном Go

Подробно: [03_design_patterns.md](./01_theory/03_design_patterns.md).

| Группа | Паттерн | В Go |
| :--- | :--- | :--- |
| Порождающие | Factory | `func NewX(...) *X` или `func NewX(...) Iface` |
| | Builder | functional options: `NewServer(WithTimeout(5))` |
| | Singleton | пакетная переменная + `sync.Once` |
| Структурные | Adapter | тип-обёртка, `http.HandlerFunc` (`func` → `interface`) |
| | Decorator | middleware над `http.Handler`, обёртка над интерфейсом |
| | Facade | пакет с лаконичным экспортируемым API |
| | Proxy | обёртка с той же сигнатурой (кэш, доступ) |
| Поведенческие | Strategy | `func`-тип или интерфейс из 1 метода (`sort.Slice`) |
| | Observer | каналы (`chan`) или срез callback-функций |
| | Iterator | `range`; `iter.Seq` (Go 1.23+) |
| | Command | `func()` как значение |
| | Chain of Responsibility | цепочка HTTP/gRPC middleware |

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
	s1 := NewServer()
	s2 := NewServer(WithTimeout(5))
	fmt.Println(s1.timeout, s2.timeout) // 30 5
	var desc SortStrategy = func(a, b int) bool { return a > b }
	s := []int{3, 1, 2}
	sort.Slice(s, func(i, j int) bool { return desc(s[i], s[j]) })
	fmt.Println(s) // [3 2 1]
}
```

## На code review и в проде

Подробно: [04_use_cases.md](./01_theory/04_use_cases.md).

| Видите | Применить |
| :--- | :--- |
| Тест требует реальную БД/HTTP | DIP: зависимость интерфейсом, в тесте фейк |
| Сквозная логика (логи, auth, метрики) | Decorator: `WithLogging(WithAuth(h))` |
| Сменные тарифы/форматы/политики | Strategy: `func`-значение или узкий интерфейс |
| Конструктор с 5+ опциональными полями | Functional Options (`Option func(*T)`) |

## Ловушки

- ⚠️ Переносить ООП-паттерны буквально (иерархии классов-стратегий, фабрики
  фабрик) — сначала проверьте, не хватит ли функции, интерфейса или композиции.
- ⚠️ Объявлять интерфейс на стороне реализации «на всякий случай», когда есть
  только одна реализация — нарушение YAGNI; объявляйте у потребителя.
- ⚠️ Ложный DRY: насильно объединённые случайно похожие фрагменты создают
  неправильную связанность между независимыми модулями.
- ⚠️ Разные реализации интерфейса по-разному ведут себя при ошибке или `nil` —
  нарушение контракта (LSP).
