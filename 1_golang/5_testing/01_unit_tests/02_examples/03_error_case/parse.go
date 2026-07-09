package parse

import (
	"fmt"
	"strconv"
)

func ParsePositive(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("bad int: %w", err)
	}
	if v <= 0 {
		return 0, fmt.Errorf("must be positive")
	}
	return v, nil
}
