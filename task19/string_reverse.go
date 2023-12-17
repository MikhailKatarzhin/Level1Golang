package main

import (
	"fmt"
)

func main() {
	text := "главрыба — абырвалг : а роза упала на лапу азора"

	// Вывод исходной строки
	fmt.Println("Source string: ", text)

	// Вывод перевернутой строки
	fmt.Println("Destination string: ", reverseString(text))
}

// Функция для переворачивания строки с учетом Unicode
func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
