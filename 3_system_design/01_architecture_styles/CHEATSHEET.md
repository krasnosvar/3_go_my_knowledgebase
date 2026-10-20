# Архитектура сервисов, сеть и API в Go — шпаргалка

Подробности — в [01_theory/](./01_theory/) и каноническом разделе
[DevOps System Design](../../../5_devops_sre_knowledgebase/14_system_design/01_architecture_and_api_styles/README.md).
⚠️ — частая ловушка.

## Раскладка Go-проекта и 12-Factor

Подробно:
[01_monolith_microservices_and_project_layout.md](./01_theory/01_monolith_microservices_and_project_layout.md).

| Директория / Фактор | Назначение в Go |
| :--- | :--- |
| `cmd/<app>/main.go` | Точка входа: чтение конфига, сборка графа зависимостей (DI), старт и graceful shutdown. |
| `internal/` | Приватные пакеты сервиса; компилятор Go запрещает их импорт извне родительского дерева. |
| `internal/<module>/internal/` | Изоляция внутренностей доменного модуля в модульном монолите от соседних модулей. |
| `CGO_ENABLED=0` | Статическая сборка одного бинарника под `scratch` / `distroless` образ. |
| Конфигурация (12-Factor) | Чтение `os.Getenv` один раз при старте в структуру `Config` с валидацией (fail-fast). |

## Clean Architecture, Hexagonal и DDD в Go

Подробно:
[02_clean_architecture_and_ddd_in_go.md](./01_theory/02_clean_architecture_and_ddd_in_go.md).

| Слой | Пакет в Go | От кого зависит |
| :--- | :--- | :--- |
| **Domain** | `internal/domain` | Ни от кого (только stdlib: `time`, `errors`). Содержит Entities и Value Objects. |
| **Use Cases** | `internal/usecase` | Только от `internal/domain`. **Сам объявляет узкие интерфейсы выходных портов.** |
| **Adapters** | `internal/adapter/{http,grpc,postgres}` | Входные адаптеры вызывают `usecase`; выходные реализуют интерфейсы `usecase` неявно. |
| **Composition Root** | `cmd/api/main.go` | Импортирует все слои, создаёт адаптеры и передаёт их в конструкторы `New...`. |

## Тюнинг сети: `http.Server` и `http.Transport`

Подробно:
[03_networking_and_api_styles_in_go.md](./01_theory/03_networking_and_api_styles_in_go.md).

| Параметр | Дефолт в stdlib | Рекомендация для продакшена |
| :--- | :--- | :--- |
| `Server.ReadHeaderTimeout` | `0` (бесконечно) | `3–5s` (защита от Slowloris) |
| `Server.ReadTimeout` / `WriteTimeout` | `0` (бесконечно) | `10–15s` |
| `Server.IdleTimeout` | `0` | `60–120s` (таймаут keep-alive) |
| `Client.Timeout` | `0` (бесконечно) | `5–10s` + `context.WithTimeout` на запрос |
| `Transport.MaxIdleConnsPerHost` | **`2` (!)** | `50–100` (иначе закрытие сокетов и `TIME_WAIT` шторм) |
| `Transport.MaxIdleConns` | `100` | `200–500` |

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	client := &http.Client{
		Timeout: 2 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     60 * time.Second,
		},
	}
	resp, err := client.Get(srv.URL)
	if err != nil {
		panic(err)
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body) // Обязательно дочитать до EOF!
		_ = resp.Body.Close()
	}()
	fmt.Println(resp.StatusCode) // 200
}
```

## REST (`encoding/json`) vs gRPC (`grpc-go`)

| Характеристика | REST + `encoding/json` | `grpc-go` + Protobuf |
| :--- | :--- | :--- |
| Сериализация | Рефлексия в рантайме, аллокации в куче | Кодогенерация без рефлексии, компактный бинарный формат |
| Соединения | Пул HTTP/1.1 keep-alive (или HTTP/2) | Мультиплексирование потоков в одном HTTP/2 соединении |
| Таймауты | Вручную через заголовки | Автоматическая передача дедлайна `context.Context` по сети |

## Ловушки

- ⚠️ Использовать `http.ListenAndServe` или `http.DefaultClient` в продакшене —
  нулевые таймауты приведут к утечке горутин и файловых дескрипторов.
- ⚠️ Оставить дефолтный `MaxIdleConnsPerHost = 2` в `http.Transport` при
  высококонкурентных вызовах одного микросервиса.
- ⚠️ Закрыть `resp.Body.Close()`, не вычитав тело (`io.Copy(io.Discard,
  resp.Body)`) — TCP-соединение не вернётся в пул keep-alive.
- ⚠️ Объявлять гигантский интерфейс репозитория в пакете `postgres` вместо
  узкого интерфейса в пакете-потребителе `usecase`.
- ⚠️ Балансировать `grpc-go` через обычный L4 `ClusterIP` в Kubernetes без
  учёта того, что HTTP/2 держит одно долгоживущее TCP-соединение.
