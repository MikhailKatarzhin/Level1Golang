package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "главрыба — абырвалг : а роза упала на лапу азора"

	// Вывод исходной строки
	fmt.Println("Source string: ", text)

	// Вывод перевернутой строки
	fmt.Println("Destination string: ", reverseWordSequence(text))
}

// Функция для переворачивания строки с учетом Unicode
func reverseWordSequence(input string) string {
	words := strings.Fields(input)
	length := len(words)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
