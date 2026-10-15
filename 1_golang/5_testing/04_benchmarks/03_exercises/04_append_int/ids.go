// Package ids — код для упражнения 04.
package ids

import (
	"fmt"
	"strings"
)

// FormatIDs возвращает id через запятую: "1,22,333".
func FormatIDs(ids []int) string {
	// TODO: один []byte + strconv.AppendInt вместо fmt.Sprint и strings.Join.
	parts := make([]string, 0, len(ids))
	for _, id := range ids {
		parts = append(parts, fmt.Sprint(id))
	}
	return strings.Join(parts, ",")
}
