package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	//numbers := []int{2, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10}
	numbers := []int{2, 4, 6, 8, 10}
	calculateArraySquares(numbers)
	fmt.Println("Время выполнения: ", time.Since(start).Microseconds())
}

// calculateArraySquares - Функция для расчета квадратов чисел из переданного списка с применением wait group
func calculateArraySquares(numbers []int) {
	var wg sync.WaitGroup

	// Итерация по каждому элементу в списке и запуск горутины для вычислений
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
}

// Функция для вычисления и вывода в std квадрата числа
func calculateSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := num * num
	// Вывод происходит напрямую в функции вычисления, однако так делать можно только, когда известно, что необходим просто вывод
	// В общих случаях так делать нельзя, так как это нарушает принцип единственной ответственности
	// Помимо нарушения принципа, в общем случае может требоваться не только вывод, но и какая-либо обработка информации
	// Однако такой способ позволяет экономить на передаче данных для вывода в другие процессы
	// В иных реализациях этот пункт будет вынесен отдельно
	fmt.Println(square)
}
