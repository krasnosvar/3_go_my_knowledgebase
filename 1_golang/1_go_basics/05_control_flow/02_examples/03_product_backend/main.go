package main

import (
	"fmt"
)

type CreditCard struct {
	CardNumber string
	Holder     string
}

type PayPal struct {
	Email string
}

type Crypto struct {
	WalletAddress string
}

// ProcessPayment обрабатывает платеж и форматирует чек в зависимости от типа оплаты (Type Switch)
func ProcessPayment(paymentMethod interface{}) string {
	switch method := paymentMethod.(type) {
	case CreditCard:
		return fmt.Sprintf("Оплата картой %s (Владелец: %s)", method.CardNumber, method.Holder)
	case PayPal:
		return fmt.Sprintf("Оплата через PayPal (%s)", method.Email)
	case Crypto:
		return fmt.Sprintf("Оплата в крипте на кошелек %s...", method.WalletAddress[:8])
	default:
		return "Неизвестный способ оплаты"
	}
}

func main() {
	payments := []interface{}{
		CreditCard{CardNumber: "**** **** **** 4321", Holder: "Ivan Ivanov"},
		PayPal{Email: "ivan@example.com"},
		Crypto{WalletAddress: "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
		"Просто текст (ошибка)",
	}

	fmt.Println("--- Бэкенд: Обработка платежей через Type Switch ---")

	for _, p := range payments {
		receipt := ProcessPayment(p)
		fmt.Printf("Чек: %s\n", receipt)
	}
}
