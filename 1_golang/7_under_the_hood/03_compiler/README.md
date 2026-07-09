# Compiler, modules, project layout

Раздел про инструменты сборки и структуру Go-проекта.

## Что внутри

- `01_go_src_file_structure.md` — что внутри Go-файла: package, imports, top-level декларации, init, видимость.
- `02_modules.md` — `go.mod`, `go.sum`, версионирование, replace, workspace, vendor.

## Связанные темы

- `4_standard_library` — пакеты stdlib, которые часто используются вместе.
- `7_under_the_hood/01_runtime` — что компилятор делает с твоим кодом.

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
