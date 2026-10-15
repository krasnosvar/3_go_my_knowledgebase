// Package concat — код для упражнения 02.
package concat

// Join склеивает parts через sep.
func Join(parts []string, sep string) string {
	// TODO: замените += на strings.Builder с Grow.
	s := ""
	for i, p := range parts {
		if i > 0 {
			s += sep
		}
		s += p
	}
	return s
}
