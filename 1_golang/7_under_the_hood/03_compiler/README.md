# 7.3 Compiler, modules, project layout

Инструменты сборки и структура Go-проекта. Раздел теоретический: примеров
и упражнений нет, команды из теории проверяйте у себя в терминале.

## Зачем

Чтобы понимать, откуда берутся `go build`, стандартная библиотека и
runtime в бинарнике, как работают `go.mod`/`go.sum`, `replace`,
workspace и vendor, и как разложить код проекта по пакетам.

## Пререквизиты

- установленный Go (см. раздел «0. Подготовка» в [корневом README](../../../README.md));
- базовый опыт `go run` / `go build` из [1_go_basics](../../1_go_basics/).

## Что внутри

- [01_go_src_file_structure.md](./01_go_src_file_structure.md) — структура установки Go (`GOROOT`): `bin/`, `src/`, `pkg/`, переменные окружения, процесс компиляции, что попадает в бинарник.
- [02_modules.md](./02_modules.md) — модули: `go.mod`, `go.sum`, версионирование, `replace`, workspace, vendor.

## Как проходить

1. Прочитайте [01_go_src_file_structure.md](./01_go_src_file_structure.md) и найдите описанные каталоги у себя: `ls "$(go env GOROOT)"`.
2. Прочитайте [02_modules.md](./02_modules.md) и посмотрите на `go.mod` в корне этого репозитория — это один модуль, в который входят все примеры и упражнения.

## Связанные темы

- [4_standard_library](../../4_standard_library/) — пакеты stdlib, которые часто используются вместе.
- [01_runtime](../01_runtime/) — что runtime делает с твоим кодом.

## Дальше

[04_profiling](../04_profiling/) — профилирование.

## Команды-минимум

```
go mod init <module>
go mod tidy           # привести go.mod в порядок
go mod vendor         # положить зависимости в vendor/
go build ./...        # собрать всё
go run ./cmd/app      # запустить
go test ./...         # тесты
go vet ./...          # статический анализ
gofmt -w .            # форматирование
goimports -w .        # форматирование + импорты
```

## Layout

Стандарт сообщества (https://github.com/golang-standards/project-layout):

```
.
├── cmd/
│   └── app/main.go
├── internal/
│   └── ...
├── pkg/
│   └── ...
├── api/
├── configs/
├── deployments/
├── go.mod
└── go.sum
```

`internal/` — пакеты, недоступные другим модулям. `pkg/` — публичные библиотеки. `cmd/` — точки входа.

Не каждому проекту нужен полный layout — для маленького CLI достаточно `main.go` + `go.mod`.
