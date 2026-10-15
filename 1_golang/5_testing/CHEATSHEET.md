# Testing — шпаргалка

Общая шпаргалка раздела: по `##`-секции на тему, теория — по ссылкам «Подробно».
⚠️ — частая ловушка. Версии указаны там, где API появилось недавно.

## 01. Unit-тесты

Подробно: [01_basics.md](./01_unit_tests/01_theory/01_basics.md),
[02_arrange_act_assert.md](./01_unit_tests/01_theory/02_arrange_act_assert.md),
[03_pitfalls_and_coverage.md](./01_unit_tests/01_theory/03_pitfalls_and_coverage.md).

| Правило | Как |
| :--- | :--- |
| Файл | `xxx_test.go` рядом с кодом; `package xxx` (white-box) или `package xxx_test` (black-box) |
| Тест | `func TestXxx(t *testing.T)`; подкатегория — `TestAdd_Negative` |
| Структура | Arrange → Act → Assert, фазы разделены пустой строкой |
| Сообщение | `t.Fatalf("Add(2,3) = %d; want %d", got, want)` — вход, факт, ожидание |

| Команда `go test` | Что делает |
| :--- | :--- |
| `go test ./...` | все пакеты модуля |
| `-v` | подробный вывод (`=== RUN`, `--- PASS`) |
| `-run 'TestAdd'` / `-run 'TestAdd/zero'` | фильтр по regexp имени / подтесту |
| `-race` | детектор гонок (обязателен для конкурентного кода и `t.Parallel`) |
| `-count=1` | отключить кэш результатов; `-count=100` — гонять flaky-тест |
| `-short` | `testing.Short() == true` — пропустить долгие тесты |
| `-tags=integration` | включить файлы с `//go:build integration` |
| `-cover` / `-coverprofile=cover.out` | покрытие / файл профиля → `go tool cover -html=cover.out` |
| `-bench=.` `-benchmem` `-benchtime=10s` | бенчмарки, аллокации, время на бенч |
| `-cpuprofile=cpu.out` | CPU-профиль бенча → `go tool pprof cpu.out` |

| Метод `t` | Назначение |
| :--- | :--- |
| `t.Error` / `t.Errorf` | отметить провал и **продолжить** |
| `t.Fatal` / `t.Fatalf` | отметить провал и **остановить** тест (только из горутины теста) |
| `t.Log` / `t.Logf` | лог (виден при `-v` или провале) |
| `t.Skip("причина")` | пропустить |
| `t.Helper()` | строка ошибки — в вызывающем тесте, а не в хелпере |
| `t.Cleanup(func() {...})` | teardown после теста, LIFO, срабатывает и после `Fatal` |
| `t.Parallel()` | параллельно с другими такими тестами |
| `t.Run(name, func(t *testing.T) {...})` | подтест |
| `t.TempDir()` | временный каталог, удаляется сам |
| `t.Setenv(k, v)` | env на время теста; ⚠️ panic вместе с `t.Parallel()` |
| `t.Context()` | контекст, отменяется перед `Cleanup` (Go 1.24+) |
| `t.Chdir(dir)` | сменить рабочий каталог на время теста (Go 1.24+) |

| Сравнить | Чем |
| :--- | :--- |
| Слайсы / map | `slices.Equal(a, b)` / `maps.Equal(m1, m2)` (Go 1.21+); `==` не компилируется |
| Произвольные структуры | `reflect.DeepEqual(want, got)` (медленно); `cmp.Diff(want, got)` из `github.com/google/go-cmp` — лучший diff |
| Ошибки | `errors.Is(err, ErrNotFound)` / `errors.As`, не `err.Error() == "..."` |

## 02. Table-driven тесты

Подробно: [01_pattern.md](./02_table_driven_tests/01_theory/01_pattern.md),
[02_pitfalls.md](./02_table_driven_tests/01_theory/02_pitfalls.md).

```go
func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    int
		wantErr error
	}{
		{"ok", "42", 42, nil},
		{"empty", "", 0, ErrEmpty},
		{"not number", "abc", 0, ErrFormat}, // подтест TestParse/not_number
	}
	for _, tt := range tests { // Go 1.22+: копия tt := tt не нужна
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Parse(tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Parse(%q) err = %v; want %v", tt.in, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Parse(%q) = %d; want %d", tt.in, got, tt.want)
			}
		})
	}
}
```

| Когда | Решение |
| :--- | :--- |
| Много похожих кейсов (типичный, граничный, ошибочный) | таблица + `t.Run` |
| 1–2 кейса | два теста с понятными именами |
| У каждого кейса уникальная подготовка | отдельные `Test*` |
| Структура кейса > 10 полей | вынести setup в helper или разбить таблицу |
| Полезные поля | `name`, входы, `want`, `wantErr error`, `setup func(t *testing.T)` |

## 03. Подтесты и setup

Подробно: [01_basics.md](./03_subtests/01_theory/01_basics.md),
[02_parallel_and_setup.md](./03_subtests/01_theory/02_parallel_and_setup.md).

| Задача | Код |
| :--- | :--- |
| Подтест | `t.Run("server", func(t *testing.T) {...})` → `TestStatus/server` |
| Запуск одного | `go test -run 'TestStatus/server'`; пробелы в имени → `_` |
| Параллельные подтесты | `t.Parallel()` первой строкой внутри `t.Run`; родитель ждёт их завершения |
| Setup на весь пакет | `func TestMain(m *testing.M) { setup(); code := m.Run(); teardown(); os.Exit(code) }` |
| Teardown теста/подтеста | `t.Cleanup(fn)` — после конкретного подтеста, не после всего `Test*` |
| Временные каталог / env | `t.TempDir()` / `t.Setenv("APP_ENV", "test")` |

| Где | Что выбрать |
| :--- | :--- |
| Долгий общий setup (контейнер БД) | `TestMain` |
| Ресурс одного теста | `t.Cleanup` / `t.TempDir` / `t.Setenv` |
| Хелпер подготовки | `func newSvc(t *testing.T) *Svc { t.Helper(); ... }` |

## 04. Бенчмарки

Подробно: [01_basics.md](./04_benchmarks/01_theory/01_basics.md),
[02_interpretation.md](./04_benchmarks/01_theory/02_interpretation.md),
[03_alloc_optimization.md](./04_benchmarks/01_theory/03_alloc_optimization.md).

| Задача | Код |
| :--- | :--- |
| Бенч (Go 1.24+) | `func BenchmarkParse(b *testing.B) { for b.Loop() { Parse("12345") } }` — setup до цикла не учитывается, вызов не выкидывается компилятором |
| Бенч (до 1.24) | `for i := 0; i < b.N; i++ { _ = Parse(s) }` — результат в `_`, `b.N` подбирает рантайм |
| Исключить setup (`b.N`-стиль) | `b.ResetTimer()` после подготовки; секции — `b.StopTimer()` / `b.StartTimer()` |
| Варианты входа | `b.Run("large", func(b *testing.B) {...})` |
| Масштаб по горутинам | `b.RunParallel(func(pb *testing.PB) { for pb.Next() {...} })` |
| Аллокации без флага | `b.ReportAllocs()` |
| Запуск без тестов | `go test -run '^$' -bench . -benchmem` |
| Сравнение до/после | `go test -bench=. -count=10 > before.txt`; `benchstat before.txt after.txt` (`golang.org/x/perf/cmd/benchstat`) |
| Escape analysis | `go build -gcflags="-m" ./...` → `moved to heap: x` |

| Поле вывода `BenchmarkParse-8  1234567  96.4 ns/op  24 B/op  1 allocs/op` | Значение |
| :--- | :--- |
| `-8` | GOMAXPROCS |
| `1234567` | итераций (`b.N`) |
| `ns/op` / `B/op` / `allocs/op` | время / байт / аллокаций на операцию |
| `p < 0.05` в `benchstat` | изменение значимо, не шум |

| Источник аллокаций | Замена |
| :--- | :--- |
| `fmt.Sprintf("user_%d", id)` | `"user_" + strconv.Itoa(id)`, `strings.Builder` |
| `s += w` в цикле | `strings.Builder` |
| `append` без ёмкости | `make([]T, 0, n)` |
| `[]byte(s)` / `string(b)` | работать в одной форме |
| Короткоживущие буферы | `sync.Pool` (`Get` → `Reset` → `Put`); GC может очистить пул |
| Значения через `any` | generics |

## 05. Моки и fakes

Подробно: [01_principles.md](./05_mocks/01_theory/01_principles.md),
[02_handmade_fakes.md](./05_mocks/01_theory/02_handmade_fakes.md),
[03_anti_patterns.md](./05_mocks/01_theory/03_anti_patterns.md).

```go
type UserRepo interface { // интерфейс у потребителя, 1–3 метода
	Get(ctx context.Context, id int) (User, error)
}

type fakeRepo struct { // stub: фиксированный ответ
	user User
	err  error
}

func (r fakeRepo) Get(context.Context, int) (User, error) { return r.user, r.err }

func TestGreet(t *testing.T) {
	svc := &Service{repo: fakeRepo{user: User{Name: "den"}}}
	if got, _ := svc.Greet(t.Context(), 1); got != "hello den" { // t.Context — Go 1.24+
		t.Fatalf("got %q", got)
	}
}
```

| Вид двойника | Что делает | Пример |
| :--- | :--- | :--- |
| Stub | фиксированный ответ | `stubClock{}.Now()` → `2026-01-01` |
| Fake | упрощённая рабочая реализация | `fakeRepo{items: map[int]User{}}` |
| Fake с ошибкой | первый вызов падает | поле `nextErr error`, сбросить после возврата |
| Spy | записывает аргументы | `s.sent = append(s.sent, to)` → `len(s.sent)` |
| Mock-фреймворк | проверка вызовов/порядка | `gomock`, `mockery` — только для сложных случаев |

| Мокать | Не мокать |
| :--- | :--- |
| БД, HTTP, файлы, время, рандом | stdlib (`fmt`, `strings`), чистые функции |
| Сложно воспроизводимые ошибки (timeout, сеть) | приватные функции — тестировать публичный метод |
| — | SQL-запросы — нужен integration test с настоящей БД |

| Время в тестах | Код |
| :--- | :--- |
| Интерфейс | `type Clock interface { Now() time.Time }` в конструкторе |
| Одна точка | `var now = time.Now`; в тесте подменить + вернуть в `t.Cleanup` |

## 06. Интеграционные тесты

Подробно: [01_basics.md](./06_integration_tests/01_theory/01_basics.md),
[02_setup_isolation.md](./06_integration_tests/01_theory/02_setup_isolation.md),
[03_flakiness.md](./06_integration_tests/01_theory/03_flakiness.md).

| Задача | Код |
| :--- | :--- |
| Handler без сети | `rr := httptest.NewRecorder(); h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/health", nil)); rr.Code` |
| Настоящий локальный сервер | `srv := httptest.NewServer(h); defer srv.Close(); client.Get(srv.URL)` |
| Отделить тегом | первая строка файла `//go:build integration` → `go test -tags=integration ./...` |
| Пропуск в быстром режиме | `if testing.Short() { t.Skip("integration") }` → `go test -short` |
| БД в контейнере | `testcontainers-go` + модуль `postgres`; `t.Cleanup(func() { _ = pg.Terminate(ctx) })` |
| Схема как в проде | миграции (`golang-migrate`, `goose`) перед тестами |
| Дождаться события | `select { case <-done: case <-time.After(2 * time.Second): t.Fatal("timeout") }` |
| Eventually | цикл до дедлайна: `if check() { return }; time.Sleep(10 * time.Millisecond)` |

| Изоляция БД | Скорость | Ограничение |
| :--- | :--- | :--- |
| Контейнер на тест | медленно | — |
| Контейнер на пакет (`TestMain`) + схема на тест | быстро | `DROP SCHEMA ... CASCADE` в `Cleanup` |
| `TRUNCATE ... RESTART IDENTITY CASCADE` в `Cleanup` | быстро | ⚠️ не с `t.Parallel()` |
| Транзакция + `Rollback` в `Cleanup` | очень быстро | ⚠️ не работает, если код сам открывает транзакции |

| Источник flaky | Лечение |
| :--- | :--- |
| `time.Sleep` как синхронизация | каналы, `sync.WaitGroup`, `select` с таймаутом |
| Жёсткие пороги времени | запас или подмена часов |
| Внешняя сеть | `httptest.NewServer` |
| Общий ресурс (файл, БД, глобалка) | изоляция, `t.TempDir` |
| Случайные данные | фиксированный seed |
| Зависимость от порядка | независимые тесты; проверка `-count=100 -race` |

## Ловушки

- ⚠️ Тест «прошёл» из кэша и не запускался — `-count=1`.
- ⚠️ `t.Fatal` из чужой горутины не останавливает тест — передавайте ошибку в канал или используйте `t.Error`.
- ⚠️ До Go 1.22 без `tt := tt` параллельные подтесты видят последний `tt`.
- ⚠️ `t.Setenv` + `t.Parallel()` в одном тесте или у предка — panic.
- ⚠️ `time.Now()` в инициализации таблицы вычисляется один раз — переносите внутрь `t.Run`.
- ⚠️ `got, _ := f()` в тесте — ветки ошибок не проверены.
- ⚠️ Глобальное состояние между тестами — зависимость от порядка, с `t.Parallel()` — гонка.
- ⚠️ `TestMain`: `os.Exit` не выполняет `defer` — teardown до `os.Exit(code)`.
- ⚠️ Бенч в стиле `b.N` без использования результата — компилятор может выкинуть вызов; `b.Loop()` (1.24+) от этого защищает.
- ⚠️ Один прогон бенча ничего не значит — `-count >= 5` и `benchstat`; бенчи разных версий Go не сравнивают.
- ⚠️ Покрытие — индикатор, не цель: 100 % без assert ничего не гарантирует; ориентир 60–80 %.
- ⚠️ Проверка «метод вызван ровно N раз» — тест деталей реализации; проверяйте поведение.
- ⚠️ Глобальная подмена (`now = ...`) без `t.Cleanup` — следующий тест получит ваш мок.
- ⚠️ Flaky-тест чинят поиском причины, а не ретраем в CI.
