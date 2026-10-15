package main

import (
	"fmt"
)

// multiplyByPowerOfTwo умножает val на 2^shift с помощью сдвига влево
func multiplyByPowerOfTwo(val int, shift uint) int {
	return val << shift
}

// divideByPowerOfTwo делит val на 2^shift с помощью сдвига вправо
func divideByPowerOfTwo(val int, shift uint) int {
	return val >> shift
}

func main() {
	fmt.Println("--- Проверка 05_bit_shift (Решение) ---")

	m := multiplyByPowerOfTwo(5, 3) // 5 * 8 = 40
	d := divideByPowerOfTwo(40, 2)  // 40 / 4 = 10

	if m == 40 && d == 10 {
		fmt.Println("Тест bit_shift: PASS")
	} else {
		fmt.Printf("Тест bit_shift: FAIL (multiply=%d, divide=%d)\n", m, d)
	}
}
