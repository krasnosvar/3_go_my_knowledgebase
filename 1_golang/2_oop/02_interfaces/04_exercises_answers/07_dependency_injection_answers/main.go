// Упражнение 07: внедрение зависимости через интерфейс (Решение).
package main

import (
	"errors"
	"fmt"
)

// Sender отправляет сообщение.
type Sender interface {
	Send(msg string) error
}

// Service — бизнес-логика уведомлений.
type Service struct {
	sender Sender
}

// recordingSender — тестовая реализация: запоминает отправленное.
type recordingSender struct {
	sent []string
	err  error
}

func (r *recordingSender) Send(msg string) error {
	if r.err != nil {
		return r.err
	}
	r.sent = append(r.sent, msg)
	return nil
}

func (s Service) Notify(user, msg string) error {
	if msg == "" {
		return errors.New("empty message")
	}
	return s.sender.Send(user + ": " + msg)
}

func main() {
	fmt.Println("--- Проверка 07_dependency_injection (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	rec := &recordingSender{}
	svc := Service{sender: rec}

	err := svc.Notify("den", "hello")
	check(err == nil, fmt.Sprintf("Notify: ожидалось без ошибки, получено %v", err))
	check(len(rec.sent) == 1 && rec.sent[0] == "den: hello", fmt.Sprintf("отправлено: получено %q", rec.sent))

	err = svc.Notify("den", "")
	check(err != nil && err.Error() == "empty message", fmt.Sprintf("пустое сообщение: ожидалась ошибка \"empty message\", получено %v", err))
	check(len(rec.sent) == 1, "пустое сообщение не должно отправляться")

	failing := Service{sender: &recordingSender{err: errors.New("smtp down")}}
	err = failing.Notify("den", "hi")
	check(err != nil && err.Error() == "smtp down", fmt.Sprintf("ошибка Sender должна пробрасываться, получено %v", err))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
