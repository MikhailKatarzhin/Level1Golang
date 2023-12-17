package main

import "fmt"

// Обмен значений между переменными, используя механизм возврата нескольких значений в Go
func main() {
	val1 := 1
	val2 := 2

	fmt.Printf("Before swap: %d, %d", val1, val2)

	val1, val2 = val2, val1

	fmt.Printf("\nAfter swap: %d, %d", val1, val2)
}
