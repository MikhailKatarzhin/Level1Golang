package main

import (
	"Level1Golang/task01"
)

func main() {
	// Используем конструкторы для создания экземпляров
	act := task01.NewAction("Иван", 1, "Е~еху!")

	// Вызываем метод из структуры Action
	act.DoWhoop() // Метод DoWhoop из Action

	println()
	// НЕ Используем "приватное" поле и используем "публичное" из встраиваемой структуры Human
	// println(act.name)       // Невозможность доступа к "приватному" полю из встроенного Human

	// Вызываем метод из встраиваемой структуры Human
	act.IntroduceYourself() // Метод IntroduceYourself из встроенного Human

	println()
	// Используем замещённый метод
	println(act.GiveYourselfName())

	// Используем одноимённый метод встроенной в Action структуры Human
	println(act.Human.GiveYourselfName())
}
