package main

import (
	"fmt"
	"github.com/MikhailKatarzhin/Level1Golang/task24/point"
)

func main() {
	// Создаем две точки с помощью конструктора
	point1 := point.NewPoint(0, 0)
	point2 := point.NewPoint(3, 4)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)
	fmt.Printf("Distance beetwen points: %.2f\n", distance)
}
