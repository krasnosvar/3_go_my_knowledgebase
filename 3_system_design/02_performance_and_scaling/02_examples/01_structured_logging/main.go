// Пример 01: структурное логирование через log/slog (stdlib).
//
// Настраиваем JSON-логгер с уровнем INFO, пишем записи с полями и
// разными уровнями, создаём дочерний логгер с request_id через With
// (корреляция) — и программа завершается.
//
// Демонстрирует: slog.New, slog.NewJSONHandler, уровни, атрибуты,
// slog.Logger.With, slog.Group, отбрасывание Debug ниже порога.
//
// Вывод детерминированный: время фиксируем через ReplaceAttr, чтобы
// метка времени не «плавала» от запуска к запуску.
//
// Запуск: go run ./02_examples/01_structured_logging/main.go
package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	// ReplaceAttr фиксирует поле time на постоянное значение,
	// чтобы вывод был детерминированным (для примера/тестов).
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				a.Value = slog.TimeValue(time.Date(2026, 7, 9, 12, 0, 0, 0, time.UTC))
			}
			return a
		},
	}

	// JSON-хендлер в stdout: каждая запись — одна JSON-строка.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	// Обычная запись с полями. msg постоянный, переменное — в атрибутах.
	logger.Info("service started", "port", 8080, "env", "prod")

	// DEBUG ниже порога INFO — эта строка не будет напечатана.
	logger.Debug("this debug line is dropped", "verbose", true)

	// WARN и ERROR — тоже структурно.
	logger.Warn("slow dependency", "dependency", "db", "latency_ms", 250)

	// Дочерний логгер с request_id: все его записи наследуют это поле.
	// Это базовый механизм корреляции записей одного запроса.
	reqLogger := logger.With("request_id", "req-42", "user", "alice")
	reqLogger.Info("handling request", "method", "GET", "path", "/order")
	reqLogger.Error("request failed", "status", 500, "err", "upstream timeout")

	// slog.Group группирует связанные поля во вложенный объект.
	logger.Info("http request",
		slog.Group("http",
			"method", "POST",
			"path", "/checkout",
			"status", 200,
		),
	)
}
