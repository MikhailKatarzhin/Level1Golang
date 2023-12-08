package main

import (
	"github.com/MikhailKatarzhin/Level1Golang/task01"
)

func main() {
	// Используем конструкторы для создания экземпляров
	act := task01.NewAction("Иван", 1, "Е~еху!")

	// Вызываем метод из структуры Action
	act.DoWhoop() // Метод DoWhoop из Action

	println()
	// Используем приватное поле и вызываем метод из встраиваемой структуры Human
	println(act.name)       // Доступ к приватному полю из встроенного Human
	act.IntroduceYourself() // Метод IntroduceYourself из встроенного Human

	println()
	// Используем замещённый метод
	println(act.giveYourselfName())

	// Используем одноимённый метод встроенной в Action структуры Human
	println(act.Human.giveYourselfName())
}
