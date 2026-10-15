// Package fizzbuzz — код под тестом для упражнения 04.
package fizzbuzz

import "strconv"

// FizzBuzz: "Fizz" для кратных 3, "Buzz" для кратных 5, "FizzBuzz" для
// кратных 15, иначе само число.
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(n)
}
