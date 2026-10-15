// Упражнение 02: внедрение ошибки через fake.
//
// Задача: напишите fakePayments с полем err error — Charge возвращает его.
// Тесты:
//   - TestPlaceOK: err == nil -> Place возвращает nil;
//   - TestPlaceChargeFails: fake возвращает errDeclined; Place должен
//     вернуть ошибку, для которой errors.Is(err, errDeclined) == true и
//     текст которой содержит id заказа (strings.Contains).
//
// Так проверяется путь ошибки, который с реальным сервисом трудно воспроизвести.
//
// Запуск: go test ./03_exercises/02_error_injection/
package order

import "testing"

func TestPlace(t *testing.T) {
	t.Fatal("TODO: fake с внедряемой ошибкой")
}
