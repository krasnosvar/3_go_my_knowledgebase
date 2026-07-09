// Пример 03: порты и адаптеры (hexagonal).
//
// Один порт — интерфейс Notifier — и два адаптера, которые его реализуют:
// ConsoleNotifier (печатает в переданный io.Writer) и EmailNotifier
// («email»-заглушка, тоже пишет в writer, имитируя отправку письма).
// Выбор адаптера делается в main.
//
// Демонстрирует: порт как точка расширения, взаимозаменяемые адаптеры,
// потребитель (usecase) зависит только от порта.
package main

import (
	"bytes"
	"fmt"
	"io"
)

// Notifier — порт: контракт «отправить уведомление».
type Notifier interface {
	Notify(to, message string) error
}

// ConsoleNotifier — адаптер: выводит уведомление как консольную строку.
type ConsoleNotifier struct {
	out io.Writer
}

func NewConsoleNotifier(out io.Writer) *ConsoleNotifier {
	return &ConsoleNotifier{out: out}
}

func (c *ConsoleNotifier) Notify(to, message string) error {
	_, err := fmt.Fprintf(c.out, "[console] %s: %s\n", to, message)
	return err
}

// EmailNotifier — адаптер-заглушка: имитирует отправку email.
// В реальности здесь был бы SMTP-клиент; для детерминизма пишем в writer.
type EmailNotifier struct {
	out  io.Writer
	from string
}

func NewEmailNotifier(out io.Writer, from string) *EmailNotifier {
	return &EmailNotifier{out: out, from: from}
}

func (e *EmailNotifier) Notify(to, message string) error {
	_, err := fmt.Fprintf(e.out, "[email] from=%s to=%s subject=Notification body=%q\n", e.from, to, message)
	return err
}

// AlertUsecase — потребитель порта. Про конкретный адаптер не знает.
type AlertUsecase struct {
	notifier Notifier
}

func NewAlertUsecase(n Notifier) *AlertUsecase {
	return &AlertUsecase{notifier: n}
}

func (a *AlertUsecase) Raise(user string) error {
	return a.notifier.Notify(user, "система требует внимания")
}

func main() {
	fmt.Println("--- Порты и адаптеры: один Notifier, два адаптера ---")

	var buf bytes.Buffer

	// Выбор адаптера в main. Меняя реализацию порта, меняем канал доставки —
	// сам usecase остаётся неизменным.
	adapters := []Notifier{
		NewConsoleNotifier(&buf),
		NewEmailNotifier(&buf, "noreply@example.com"),
	}

	for _, adapter := range adapters {
		uc := NewAlertUsecase(adapter)
		if err := uc.Raise("user@example.com"); err != nil {
			fmt.Printf("ошибка уведомления: %v\n", err)
			return
		}
	}

	// Печатаем накопленный вывод обоих адаптеров.
	fmt.Print(buf.String())
	fmt.Println("Готово, программа завершается.")
}
