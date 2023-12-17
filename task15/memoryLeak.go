package main

import (
	"strings"
)

var justString string

func createHugeString(nSymbols int) string {
	hs := strings.Builder{}

	for i := 0; i < nSymbols; i++ {
		hs.WriteRune(' ')
	}

	return hs.String()
}

// Так как justString объявлена за пределами someFunc и является слайсом v
// то даже после завершения someFunc строка полученная в  createHugeString будет занимать память
// так как justString как слайс будет на неё ссылаться
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}

func someFunc() {
	v := createHugeString(1 << 10)

	// Создаем новый срез рун нужной длины и копируем туда первые 100 рун из v
	// таким образом v останется единственным, что будет ссылаться на большую строку
	// и по окончании функции, сборщик мусора освободить память, занимаемую ею
	runes := []rune(v)[:100]

	// Создаем строку из среза рун
	justString = string(runes)

}

func main() {
	someFunc()
}
