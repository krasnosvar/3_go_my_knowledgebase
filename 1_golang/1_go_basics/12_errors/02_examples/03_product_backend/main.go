package main

import (
	"errors"
	"fmt"
)

// APIError содержит структурированные данные об ошибке для API
type APIError struct {
	Code    int
	Message string
}

// pointer receiver
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error [%d]: %s", e.Code, e.Message)
}

// performAction симулирует бизнес-логику бэкенда
func performAction(userID string) error {
	if userID == "" {
		// Ошибка валидации
		return fmt.Errorf("user authentication failed: %w", &APIError{
			Code:    401,
			Message: "Unauthorized: Invalid or missing token",
		})
	}
	return nil
}

func main() {
	fmt.Println("--- Обработчик ошибок API на основе errors.As ---")

	err := performAction("")
	if err != nil {
		// Извлекаем APIError из цепочки ошибок
		var apiErr *APIError
		if errors.As(err, &apiErr) {
			fmt.Printf("Отправляем HTTP JSON ответ:\n  {\n    \"status_code\": %d,\n    \"error\": %q\n  }\n",
				apiErr.Code, apiErr.Message)
		} else {
			// Обычная неизвестная ошибка (HTTP 500)
			fmt.Println("Отправляем HTTP JSON ответ:\n  {\n    \"status_code\": 500,\n    \"error\": \"Internal Server Error\"\n  }")
		}
	}
}
