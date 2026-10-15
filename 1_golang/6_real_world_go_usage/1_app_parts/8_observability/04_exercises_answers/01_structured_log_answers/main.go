// Упражнение 01: структурные логи через log/slog (Решение).
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

func logRequest(logger *slog.Logger, method, path string, status int, dur time.Duration) {
	level := slog.LevelInfo
	if status >= 500 {
		level = slog.LevelError
	}
	logger.Log(context.Background(), level, "request",
		"method", method,
		"path", path,
		"status", status,
		"duration_ms", dur.Milliseconds(),
	)
}

func main() {
	fmt.Println("--- Проверка 01_structured_log (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logRequest(logger, "GET", "/users", 200, 42*time.Millisecond)
	logRequest(logger, "POST", "/orders", 503, time.Second)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	check(len(lines) == 2, "две записи в логе")
	if len(lines) == 2 {
		var a, b map[string]any
		_ = json.Unmarshal([]byte(lines[0]), &a)
		_ = json.Unmarshal([]byte(lines[1]), &b)
		check(a["level"] == "INFO" && a["msg"] == "request", "200 → INFO request")
		check(a["method"] == "GET" && a["path"] == "/users", "поля method и path")
		check(a["status"] == 200.0 && a["duration_ms"] == 42.0, "status=200, duration_ms=42 числами")
		check(b["level"] == "ERROR", "503 → ERROR")
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
