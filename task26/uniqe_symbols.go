package main

import (
	"fmt"
	"strings"
)

func main() {
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range stringsToCheck {
		result := isUnique(str)
		fmt.Printf("Uniqe of string [%s]: %t\n", str, result)
	}
}

func isUnique(str string) bool {
	// Создаем карту для отслеживания символов
	charMap := make(map[rune]bool)

	// Приводим строку к нижнему регистру и
	// Проходим по всем символам в строке
	for _, char := range strings.ToLower(str) {
		// Проверяем, есть ли уже такой символ в карте
		if charMap[char] {
			return false // Если символ уже встречался, возвращаем false
		}
		// Добавляем символ в карту
		charMap[char] = true
	}
	return true // Если все символы уникальны, возвращаем true
}
