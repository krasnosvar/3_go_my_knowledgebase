# Упражнения: encoding/json, encoding/csv, encoding/base64

Каждая задача — это `main.go` с функцией-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — ваша цель довести их до `PASS`. Только стандартная библиотека, без внешней сети.

Запуск шаблона (из папки темы):

```bash
go run ./03_exercises/01_json_tags/main.go
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_json_tags_answers/main.go
```

## Список задач (от простого к сложному)

1. [`01_json_tags`](./01_json_tags/main.go) — теги `json:"id"`, `omitempty`, `json:"-"` для секрета.
2. [`02_json_decode_array`](./02_json_decode_array/main.go) — декодирование JSON-массива объектов, ошибки на битом вводе.
3. [`03_csv_parse`](./03_csv_parse/main.go) — `csv.Reader` с заголовком, кавычки внутри полей, ошибка с номером строки.
4. [`04_csv_write`](./04_csv_write/main.go) — `csv.Writer` в `bytes.Buffer`, `Flush`/`WriteAll`, экранирование.
5. [`05_base64_token`](./05_base64_token/main.go) — токен `user:exp` в `RawURLEncoding` и обратный разбор (`strings.Cut`).
6. [`06_json_validation`](./06_json_validation/main.go) — строгий разбор: `DisallowUnknownFields` + бизнес-валидация.
7. [`07_stream_decoder`](./07_stream_decoder/main.go) — потоковое чтение NDJSON через `Decoder` до `io.EOF`.
8. [`08_api_payload`](./08_api_payload/main.go) — конверт ответа `data`/`meta`, `[]` вместо `null` для пустого списка.
