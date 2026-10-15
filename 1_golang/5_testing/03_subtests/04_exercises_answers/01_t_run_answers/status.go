// Package status — код под тестом для упражнения 01.
package status

// Class возвращает класс HTTP-кода: "info", "success", "redirect",
// "client_error", "server_error" или "unknown".
func Class(code int) string {
	switch {
	case code >= 100 && code < 200:
		return "info"
	case code >= 200 && code < 300:
		return "success"
	case code >= 300 && code < 400:
		return "redirect"
	case code >= 400 && code < 500:
		return "client_error"
	case code >= 500 && code < 600:
		return "server_error"
	}
	return "unknown"
}
