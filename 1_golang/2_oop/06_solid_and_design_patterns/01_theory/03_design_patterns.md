# 03. Паттерны проектирования в Go

Паттерны GoF — типовые решения повторяющихся задач. В Go многие классические
паттерны **упрощаются или растворяются**: там, где в Java/C++ нужна иерархия
классов, в Go достаточно интерфейса, функции-значения или композиции. Ниже —
три группы с акцентом на идиоматичный Go.

## Порождающие (creational) — как создавать объекты

- **Factory / Factory Method** — функция-конструктор, возвращающая структуру
  или интерфейс. В Go это просто `func NewX(...) *X` или `func NewX(...) Iface`,
  никаких иерархий фабрик не нужно.
- **Builder** — пошаговая сборка сложного объекта. Идиоматичная альтернатива в
  Go — **functional options**: `NewServer(WithTimeout(5), WithTLS())`.
- **Singleton** — единственный экземпляр. В Go — пакетная переменная +
  `sync.Once` (или передача зависимости явно из `main`).
- **Abstract Factory / Prototype** — редки в Go; обычно избыточны.

```go
// Functional options — идиоматичная замена Builder.
type Server struct{ timeout int }
type Option func(*Server)

func WithTimeout(t int) Option { return func(s *Server) { s.timeout = t } }

func NewServer(opts ...Option) *Server {
	s := &Server{timeout: 30} // разумные значения по умолчанию
	for _, o := range opts {
		o(s)
	}
	return s
}
```

## Структурные (structural) — как собирать объекты в структуры

- **Adapter** — приводит один интерфейс к другому. В Go часто это тип-обёртка
  или `http.HandlerFunc`-подобный адаптер `func → interface`.
- **Decorator** — оборачивает объект, добавляя поведение, не меняя его.
  Идиоматично в Go: middleware над `http.Handler`, обёртка над интерфейсом
  (логирование, метрики, кэш). См. пример `03_decorator_pattern`.
- **Facade** — упрощённый фасад над сложной подсистемой; в Go — просто пакет с
  удобным экспортируемым API.
- **Proxy** — подставной объект с той же сигнатурой (кэш, ленивая загрузка,
  контроль доступа) — тоже обёртка вокруг интерфейса.
- **Composite** — дерево однотипных объектов через общий интерфейс.

```go
// Decorator: обёртка вокруг интерфейса добавляет логирование.
type Service interface{ Do(x int) int }

type LoggingService struct{ next Service }

func (l LoggingService) Do(x int) int {
	res := l.next.Do(x)
	fmt.Printf("Do(%d) = %d\n", x, res)
	return res
}
```

## Поведенческие (behavioral) — как объекты взаимодействуют

- **Strategy** — взаимозаменяемые алгоритмы. В Go — интерфейс с одним методом
  **или просто func-тип**, который подставляют. См. пример
  `02_strategy_pattern`.
- **Observer** — подписка на события. В Go часто заменяется каналами (`chan`)
  или срезом callback-функций.
- **Iterator** — обход коллекции. В Go — `range`, а с Go 1.23 —
  итераторы-функции (`iter.Seq`).
- **Command** — инкапсуляция действия; в Go — просто `func()` как значение.
- **Template Method** — каркас алгоритма с переопределяемыми шагами; в Go
  достигается композицией и передачей функций, а не наследованием.
- **State / Chain of Responsibility** — состояния и цепочки обработчиков;
  цепочки естественно ложатся на HTTP/gRPC middleware.

```go
// Strategy через func-тип — самый идиоматичный вариант в Go.
type SortStrategy func(a, b int) bool

var Ascending SortStrategy = func(a, b int) bool { return a < b }
var Descending SortStrategy = func(a, b int) bool { return a > b }
```

## Главная мысль

В Go **функции — значения первого класса**, а интерфейсы неявные и узкие.
Поэтому:

- Strategy / Command / Observer часто = функция или канал, а не отдельная
  иерархия классов.
- Decorator / Proxy / Adapter = тип-обёртка вокруг интерфейса.
- Factory / Builder = конструктор `New...` и functional options.

Не нужно «переносить» ООП-паттерны буквально: сначала спросите, не решается ли
задача композицией, интерфейсом или функцией.

> ⚠️ Частый вопрос на собеседовании: как реализовать Strategy в Go идиоматично?
> Через func-тип (алгоритм — это значение-функция, которое подставляется) или
> через маленький интерфейс с одним методом. Иерархия классов-стратегий не
> нужна.

> ⚠️ Частый вопрос на собеседовании: где в стандартной библиотеке Go видны
> паттерны? Decorator/Adapter — `http.HandlerFunc` и middleware; Strategy —
> `sort.Slice` с функцией сравнения; Singleton — `sync.Once`; Iterator —
> `range` и `iter.Seq` (Go 1.23+); Observer — каналы.
