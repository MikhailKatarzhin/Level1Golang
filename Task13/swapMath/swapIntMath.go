package main

import "fmt"

// Подход через математические вычисления может привести к переполнению
func main() {
	val1 := 1
	val2 := 2

	fmt.Printf("Before swap: %d, %d", val1, val2)

	val1 = val1 + val2
	val2 = val1 - val2
	val1 = val1 - val2

	fmt.Printf("\nAfter swap: %d, %d", val1, val2)
}
