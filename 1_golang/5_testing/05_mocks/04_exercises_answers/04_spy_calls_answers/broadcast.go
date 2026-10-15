// Package broadcast — код под тестом для упражнения 04.
package broadcast

// Notifier отправляет сообщение пользователю.
type Notifier interface {
	Send(user, msg string) error
}

// Broadcast отправляет msg всем users по порядку и останавливается на
// первой ошибке, возвращая её.
func Broadcast(n Notifier, users []string, msg string) error {
	for _, u := range users {
		if err := n.Send(u, msg); err != nil {
			return err
		}
	}
	return nil
}
