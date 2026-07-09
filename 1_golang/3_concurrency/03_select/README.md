# Изучение select в Go

`select` — многоканальный switch для каналов.

## Что важно

- ожидание нескольких каналов;
- `default` для non-blocking поведения;
- timeout pattern (`time.After`);
- cancel/stop сигналы;
- multiplexing потоков.
