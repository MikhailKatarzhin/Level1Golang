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
	var output strings.Builder

	for i := len(words) - 1; i > 0; i-- {
		output.WriteString(words[i])
		output.WriteRune(' ')
	}

	output.WriteString(words[0])

	return output.String()
}
