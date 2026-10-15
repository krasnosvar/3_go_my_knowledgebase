# Реальные use-cases интерфейсов

## 1. CLI

Паттерн command registry: разные команды реализуют общий интерфейс `Run(args []string) error`.

## 2. Инфраструктура

Абстракция хранилищ: `ConfigStore`, `SecretProvider`, `Locker`.

## 3. Backend

Repository, payment adapter, notification sender, логгер, кэш — типичные точки расширения.

## 4. Что говорить на собеседовании

- интерфейсы в Go реализуются неявно;
- маленькие интерфейсы легче поддерживать;
- pointer receiver влияет на удовлетворение интерфейса;
- `nil interface` — частая причина багов.
