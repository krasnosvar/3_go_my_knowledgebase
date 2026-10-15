// Упражнение 07: запуск внешней команды с таймаутом (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func run(timeout time.Duration, name string, args []string, env []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Env = append(os.Environ(), env...)
	out, err := cmd.Output()
	if ctx.Err() != nil {
		return "", ctx.Err()
	}
	return string(out), err
}

func main() {
	// Режим дочернего процесса.
	switch os.Getenv("CHILD_MODE") {
	case "echo":
		fmt.Print("hello")
		return
	case "sleep":
		time.Sleep(5 * time.Second)
		return
	}

	fmt.Println("--- Проверка 07_exec_timeout (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	self, err := os.Executable()
	if err != nil {
		fmt.Println("FAIL: os.Executable:", err)
		return
	}

	out, err := run(5*time.Second, self, nil, []string{"CHILD_MODE=echo"})
	check(err == nil && out == "hello", "быстрая команда → stdout \"hello\"")
	start := time.Now()
	_, err = run(200*time.Millisecond, self, nil, []string{"CHILD_MODE=sleep"})
	check(errors.Is(err, context.DeadlineExceeded), "медленная команда → DeadlineExceeded")
	check(time.Since(start) < 3*time.Second, "процесс убит по таймауту, а не дождался конца")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
