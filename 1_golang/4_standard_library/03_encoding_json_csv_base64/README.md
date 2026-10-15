# 03. encoding/json, encoding/csv, encoding/base64

Почти каждый backend-сервис принимает и отдаёт JSON, импортирует/экспортирует CSV и передаёт бинарные данные в текстовом виде через base64. Тема разбирает struct-теги, строгий разбор запросов, потоковый `json.Decoder`, корректное экранирование CSV и разницу между вариантами base64.

## 🧭 Пререквизиты

- [1_go_basics/08_structs](../../1_go_basics/08_structs/) — структуры и теги
- [1_go_basics/12_errors](../../1_go_basics/12_errors/) — обработка ошибок
- [2_oop/02_interfaces](../../2_oop/02_interfaces/) — `io.Reader`/`io.Writer` как интерфейсы
- [other/05_encoding_hashing_encryption.md](../../../other/05_encoding_hashing_encryption.md) и [other/06_serialization_formats.md](../../../other/06_serialization_formats.md) — кодирование vs шифрование, форматы сериализации

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_json.md](./01_theory/01_json.md) — `encoding/json`: `Marshal`/`Unmarshal`, `Encoder`/`Decoder`, теги, кастомизация, `any`.
    *   [02_csv.md](./01_theory/02_csv.md) — `encoding/csv`: чтение, запись, кавычки и экранирование, заголовок, частые ошибки.
    *   [03_base64.md](./01_theory/03_base64.md) — `encoding/base64`: Std/URL/Raw-варианты, размер, use-cases, hex как альтернатива.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_json_roundtrip](./02_examples/01_json_roundtrip/main.go) — JSON туда и обратно.
    *   [02_csv_read](./02_examples/02_csv_read/main.go) — чтение CSV.
    *   [03_base64_encode_decode](./02_examples/03_base64_encode_decode/main.go) — кодирование и декодирование base64.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_json_tags](./03_exercises/01_json_tags/main.go) — теги `json:"id"`, `omitempty`, `json:"-"` для секрета.
    *   [02_json_decode_array](./03_exercises/02_json_decode_array/main.go) — декодирование JSON-массива объектов, ошибки на битом вводе.
    *   [03_csv_parse](./03_exercises/03_csv_parse/main.go) — `csv.Reader` с заголовком, кавычки внутри полей, ошибка с номером строки.
    *   [04_csv_write](./03_exercises/04_csv_write/main.go) — `csv.Writer` в `bytes.Buffer`, `Flush`/`WriteAll`, экранирование.
    *   [05_base64_token](./03_exercises/05_base64_token/main.go) — токен `user:exp` в `RawURLEncoding` и обратный разбор (`strings.Cut`).
    *   [06_json_validation](./03_exercises/06_json_validation/main.go) — строгий разбор: `DisallowUnknownFields` + бизнес-валидация.
    *   [07_stream_decoder](./03_exercises/07_stream_decoder/main.go) — потоковое чтение NDJSON через `Decoder` до `io.EOF`.
    *   [08_api_payload](./03_exercises/08_api_payload/main.go) — конверт ответа `data`/`meta`, `[]` вместо `null` для пустого списка.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Команды выполняются из папки темы (`1_golang/4_standard_library/03_encoding_json_csv_base64`).

1. Прочитайте теорию в [01_theory/](./01_theory/), начиная с [01_json.md](./01_theory/01_json.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_json_roundtrip/main.go
   ```
3. Решите задачи в `03_exercises/` и запустите самопроверку (в стартовом состоянии печатает `FAIL`):
   ```bash
   go run ./03_exercises/01_json_tags/main.go
   ```
4. Сверьтесь с эталоном:
   ```bash
   go run ./04_exercises_answers/01_json_tags_answers/main.go
   ```

## ➡️ Дальше

[04_io_bufio_os_filepath](../04_io_bufio_os_filepath/) — потоки, файлы и пути.
