# Разное, но важное (Computer Science Basics)

Небольшие пояснительные заметки по общим основам, которые полезно понимать
бэкенд-разработчику и часто всплывают на собеседованиях, но не относятся к
какому-то одному разделу базы (языку, БД или system design).

Каждый файл — самостоятельная короткая тема: можно читать в любом порядке.

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — общая шпаргалка по всем заметкам: кодировки, числа, endianness, хеши и шифрование, форматы, сборка, сигналы, время.

## 🗂 Темы

1. [01_binary_vs_text_files.md](./01_binary_vs_text_files.md) — бинарные vs текстовые файлы: почему это вопрос интерпретации байтов, а не физического хранения.
2. [02_encodings_ascii_unicode_utf8.md](./02_encodings_ascii_unicode_utf8.md) — кодировки символов: ASCII, Unicode (code points), UTF-8/16/32; как это связано с `string` и `rune` в Go.
3. [03_number_representation.md](./03_number_representation.md) — как хранятся числа: дополнительный код, переполнение, IEEE 754 и почему `0.1 + 0.2 != 0.3`.
4. [04_endianness.md](./04_endianness.md) — порядок байтов: big-endian vs little-endian, network byte order, `encoding/binary`.
5. [05_encoding_hashing_encryption.md](./05_encoding_hashing_encryption.md) — чем отличаются кодирование, хеширование и шифрование (частая путаница на собеседовании).
6. [06_serialization_formats.md](./06_serialization_formats.md) — сериализация: текстовые (JSON/XML/YAML) vs бинарные (Protobuf/gob) форматы, когда что.
7. [07_compilation_linking.md](./07_compilation_linking.md) — компиляция vs интерпретация, этапы сборки, линковка; особенности Go (статический бинарник, кросс-компиляция).
8. [08_processes_signals_streams.md](./08_processes_signals_streams.md) — процессы, exit codes, переменные окружения, stdin/stdout/stderr, Unix-сигналы и graceful shutdown.
9. [09_time_timezones.md](./09_time_timezones.md) — время и часовые пояса: UTC, Unix timestamp, DST, монотонные часы vs wall-clock.

## 🚀 Как пользоваться

Раздел справочный — открывайте нужную тему по мере необходимости. Многие заметки
ссылаются на соответствующие разделы базы (`1_golang`, `3_system_design`), где
тема раскрыта глубже с примерами и упражнениями.
