// Упражнение 01: gRPC status codes.
//
// В реальном сервисе используется google.golang.org/grpc/codes и status;
// здесь коды смоделированы на stdlib, чтобы сосредоточиться на смысле.
//
// Задача:
//   - toCode(err error) code: nil → OK; errors.Is на доменные ошибки:
//     errValidation → InvalidArgument, errNotFound → NotFound,
//     errExists → AlreadyExists, errNoAccess → PermissionDenied,
//     context.DeadlineExceeded → DeadlineExceeded; прочее → Internal;
//   - httpStatus(c code) int — как grpc-gateway: OK 200, InvalidArgument 400,
//     NotFound 404, AlreadyExists 409, PermissionDenied 403,
//     DeadlineExceeded 504, Internal 500.
//
// Запуск: go run ./03_exercises/01_status_codes
package main

import (
	"context"
	"errors"
	"fmt"
)

type code int

// Значения совпадают с google.golang.org/grpc/codes.
const (
	OK               code = 0
	InvalidArgument  code = 3
	DeadlineExceeded code = 4
	NotFound         code = 5
	AlreadyExists    code = 6
	PermissionDenied code = 7
	Internal         code = 13
)

var (
	errValidation = errors.New("validation")
	errNotFound   = errors.New("not found")
	errExists     = errors.New("exists")
	errNoAccess   = errors.New("no access")
)

var _ = context.DeadlineExceeded

func toCode(err error) code {
	// TODO: switch с errors.Is.
	return -1
}

func httpStatus(c code) int {
	// TODO: таблица соответствия code → HTTP.
	return 0
}

func main() {
	fmt.Println("--- Проверка 01_status_codes ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(toCode(nil) == OK, "nil → OK")
	check(toCode(fmt.Errorf("email: %w", errValidation)) == InvalidArgument, "validation → InvalidArgument")
	check(toCode(fmt.Errorf("order: %w", errNotFound)) == NotFound, "not found → NotFound")
	check(toCode(errExists) == AlreadyExists, "exists → AlreadyExists")
	check(toCode(errNoAccess) == PermissionDenied, "no access → PermissionDenied")
	check(toCode(fmt.Errorf("db: %w", context.DeadlineExceeded)) == DeadlineExceeded, "deadline → DeadlineExceeded")
	check(toCode(errors.New("boom")) == Internal, "прочее → Internal")
	check(httpStatus(NotFound) == 404 && httpStatus(InvalidArgument) == 400 && httpStatus(DeadlineExceeded) == 504, "HTTP-маппинг")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
