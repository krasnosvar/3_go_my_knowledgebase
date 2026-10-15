// Упражнение 01: gRPC status codes (Решение).
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

func toCode(err error) code {
	switch {
	case err == nil:
		return OK
	case errors.Is(err, errValidation):
		return InvalidArgument
	case errors.Is(err, errNotFound):
		return NotFound
	case errors.Is(err, errExists):
		return AlreadyExists
	case errors.Is(err, errNoAccess):
		return PermissionDenied
	case errors.Is(err, context.DeadlineExceeded):
		return DeadlineExceeded
	default:
		return Internal
	}
}

func httpStatus(c code) int {
	return map[code]int{
		OK: 200, InvalidArgument: 400, NotFound: 404, AlreadyExists: 409,
		PermissionDenied: 403, DeadlineExceeded: 504, Internal: 500,
	}[c]
}

func main() {
	fmt.Println("--- Проверка 01_status_codes (Решение) ---")

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
