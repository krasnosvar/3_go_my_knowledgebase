# Чистая архитектура, Порты и Адаптеры и DDD в Go

## О чём эта тема (краткий контекст)

Когда бизнес-логика сервиса перемешана с SQL-запросами, HTTP-заголовками и
JSON-тегами внешней очереди, любое изменение базы данных или формата API ломает
ядро приложения, а написать быстрые unit-тесты без поднятия инфраструктуры
становится невозможно. Слоистая архитектура, Гексагональная архитектура (Ports &
Adapters), Clean Architecture и тактический Domain-Driven Design (DDD) решают
эту проблему за счёт направления всех зависимостей внутрь — к бизнес-домену.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [Внутренняя архитектура сервиса: Clean Architecture, Hexagonal и DDD](../../../../5_devops_sre_knowledgebase/14_system_design/01_architecture_and_api_styles/01_theory/02_service_architecture_clean_ddd.md)
> - Базовые принципы ООП и SOLID в Go:
>   [06_solid_and_design_patterns](../../../1_golang/2_oop/06_solid_and_design_patterns/)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

## 1. Главное отличие Go: утиная типизация и владение интерфейсами

В Java или C# для реализации Dependency Inversion (DIP) пакет домена или
приложения вынужден экспортировать публичный интерфейс `OrderRepository`, а
пакет инфраструктуры пишет `class PostgresOrderRepo implements OrderRepository`.

В Go интерфейсы реализуются **неявно** (утиная типизация). Из этого следует
главная архитектурная идиома Go:

> **«Accept interfaces, return structs» — объявляйте интерфейс-порт в том
> пакете,
> который его потребляет (`usecase`), а не в пакете, который его реализует
> (`postgres`).**

### Как это выглядит в коде Go

```go
// --- Пакет internal/usecase (приложение / бизнес-сценарии) ---
package usecase

import (
	"context"
	"myservice/internal/domain"
)

// OrderSaver — выходной порт (driven port).
// Объявляется прямо здесь, у потребителя, и содержит ТОЛЬКО нужные методы (ISP).
type OrderSaver interface {
	Save(ctx context.Context, order domain.Order) error
}

type CreateOrderUseCase struct {
	saver OrderSaver
}

// Принимаем интерфейс (OrderSaver), возвращаем структуру (*CreateOrderUseCase).
func NewCreateOrderUseCase(s OrderSaver) *CreateOrderUseCase {
	return &CreateOrderUseCase{saver: s}
}
```

```go
// --- Пакет internal/adapter/postgres (инфраструктурный адаптер) ---
package postgres

import (
	"context"
	"database/sql"
	"myservice/internal/domain"
)

// OrderRepo даже НЕ импортирует пакет usecase!
// Но компилятор Go проверит в main.go, что *OrderRepo подходит под OrderSaver.
type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo { return &OrderRepo{db: db} }

func (r *OrderRepo) Save(ctx context.Context, order domain.Order) error {
	// SQL-запрос с использованием ctx
	return nil
}
```

**Что это даёт:**
1. Пакет `postgres` не зависит от `usecase`, а `usecase` не зависит от
   `postgres`. Оба зависят только от `internal/domain`.
2. Разные сценарии в `usecase` могут объявлять узкие интерфейсы на 1–2 метода
   (`OrderReader`, `OrderSaver`), а один `*postgres.OrderRepo` удовлетворит им
   всем одновременно без единого изменения.

## 2. Раскладка пакетов Clean / Hexagonal Architecture в Go

Для сервиса со сложной бизнес-логикой идиоматичная раскладка пакетов выглядит
так:

```text
internal/
├── domain/                  # Ядро: Entities, Value Objects, доменные ошибки
│   └── order.go             # 0 внешних импортов (только stdlib: time, errors)
├── usecase/                 # Сценарии использования + интерфейсы выходных портов
│   └── create_order.go      # Импортирует только internal/domain
└── adapter/                 # Внешний слой (Ports & Adapters)
    ├── http/                # Входной адаптер (Primary / Driving): net/http хендлеры
    ├── grpc/                # Входной адаптер: gRPC-сервер
    ├── postgres/            # Выходной адаптер (Secondary / Driven): SQL-репозиторий
    └── kafka/               # Выходной адаптер: публикация событий
```

См. рабочие примеры в этом разделе:
- `02_examples/01_dependency_injection/main.go`
- `02_examples/02_clean_layers/main.go`
- `02_examples/03_ports_adapters/main.go`

## 3. Тактический DDD в Go: Entity, Value Object и защита инвариантов

В Go нет конструкторов по умолчанию и приватных сеттеров уровня класса, но
пакетная видимость (строчная буква поля) идеально защищает инварианты агрегата:

```go
package domain

import "errors"

var ErrInvalidAmount = errors.New("amount must be positive")

// Money — неизменяемый Value Object (передаётся по значению).
type Money struct {
	amountCents int64
	currency    string
}

func NewMoney(cents int64, currency string) (Money, error) {
	if cents <= 0 || currency == "" {
		return Money{}, ErrInvalidAmount
	}
	return Money{amountCents: cents, currency: currency}, nil
}

func (m Money) Cents() int64     { return m.amountCents }
func (m Money) Currency() string { return m.currency }
```

### Почему не стоит смешивать доменные структуры с ORM/JSON-моделями

Если навесить на доменную структуру теги `db:"..."`, `gorm:"..."`, `json:"..."`
и сделать все поля публичными ради автобиндинга фреймворка, то любой внешний
пакет сможет записать `order.Status = "paid"` в обход бизнес-правил. В чистой
архитектуре HTTP-адаптер декодирует запрос в свой `createOrderRequest`, а
репозиторий маппит `domain.Order` в локальную структуру строки БД.

## 4. Сборка зависимостей (DI): `main.go` против `wire` и `fx`

В экосистеме Go есть три подхода к связыванию слоёв (Composition Root):

| Подход | Как работает | Когда выбирать |
| :--- | :--- | :--- |
| **Явный DI в `main.go`** | Вызов конструкторов `repo := NewRepo(db); svc := NewService(repo); srv := NewServer(svc)` руками. | **По умолчанию (90% сервисов).** Ноль магии, ошибки на этапе компиляции, переход к определению по F12 работает мгновенно. |
| **Кодогенерация (`google/wire`)** | Генерирует тот же код инициализации на этапе `go generate` до компиляции. | Крупные монолиты с сотнями зависимостей, где ручной `main.go` разросся на тысячи строк, но рефлексия в рантайме нежелательна. |
| **Рефлексия (`uber-go/fx`, `dig`)** | Строит граф зависимостей в рантайме при старте через пакет `reflect` и управляет жизненным циклом (`OnStart`/`OnStop`). | Когда в компании принят единый платформенный фреймворк поверх `fx`. Минус: ошибки графа всплывают при старте, а не при компиляции. |

## 5. Как не скатиться в «Java на Go» (борьба с оверинжинирингом)

Clean Architecture и DDD окупаются, когда есть сложные бизнес-инварианты. Если
ваш сервис — это простой CRUD поверх одной таблицы, который перекладывает JSON в
PostgreSQL без бизнес-правил, создание пяти слоёв (`Controller → UseCase →
Domain → Repository → DAO`) со стопроцентным дублированием структур на каждом
шаге нарушает KISS и YAGNI.

Для простого микросервиса в Go достаточно двух уровней: пакета `http`
(хендлеры) и пакета `store` (работа с БД), связанных узким интерфейсом у
хендлера для удобства тестирования.
