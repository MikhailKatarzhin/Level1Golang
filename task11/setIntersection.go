package main

import (
	"math/rand"
)

func main() {
	setLeft := fillSetByInt(newSet[int](), 50)
	setRight := fillSetByInt(newSet[int](), 50)

	setLeft.Print("left")
	setRight.Print("right")

	Intersection(setLeft, setRight).Print("intersect")
}

// fillSetByInt Заполняет множество N случайными значениями
func fillSetByInt(s Set[int], nValues int) Set[int] {
	for i := 0; i < nValues; i++ {
		// Создание случайного значения от 0 до 500 и добавление его во множество
		s.set[rand.Intn(500)] = struct{}{}
	}

	return s
}
