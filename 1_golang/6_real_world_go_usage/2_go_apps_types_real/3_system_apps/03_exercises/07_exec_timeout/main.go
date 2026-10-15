// Упражнение 07: запуск внешней команды с таймаутом.
//
// Задача: реализуйте run(timeout, name string, args []string, env []string) (string, error):
//   - exec.CommandContext с context.WithTimeout (не забудьте cancel);
//   - cmd.Env = append(os.Environ(), env...);
//   - вернуть stdout (cmd.Output());
//   - если сработал таймаут — вернуть ctx.Err() (context.DeadlineExceeded),
//     а не «signal: killed», чтобы вызывающий код мог отличить таймаут.
//
// Тест запускает этот же бинарник (os.Executable) в роли дочернего процесса,
// чтобы не зависеть от наличия sleep/echo в системе.
//
// Запуск: go run ./03_exercises/07_exec_timeout
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var _ = exec.CommandContext

func run(timeout time.Duration, name string, args []string, env []string) (string, error) {
	// TODO: CommandContext + Output + распознавание таймаута.
	return "", nil
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

	fmt.Println("--- Проверка 07_exec_timeout ---")

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
