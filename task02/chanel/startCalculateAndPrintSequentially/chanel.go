package main

import (
	"fmt"
	"time"
)

func main() {
	//numbers := []int{2, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10}
	numbers := []int{2, 4, 6, 8, 10}

	start := time.Now()
	calculateAndPrintArraySquares(numbers)
	fmt.Println("Время выполнения: ", time.Since(start).Microseconds())
}

// calculateAndPrintArraySquares конкурентно вычисляет квадраты чисел из массива и выводит их в порядке расчета
func calculateAndPrintArraySquares(numbers []int) {
	// Создание и отложенное закрытие канала для передачи результатов вычислений
	squareChan := make(chan int)
	defer close(squareChan)

	// Запуск из главного процесса вычисления квадратов чисел в горутинах
	startCalculateSquare(numbers, squareChan)

	// Вывод главным процессом результатов из канала
	printFromChanSequentially(len(numbers), squareChan)
}

// startCalculateSquare запускает процессы вычисления квадратов чисел в горутинах
func startCalculateSquare(numbers []int, squareChan chan int) {
	for _, num := range numbers {
		go calculateSquare(num, squareChan)
	}
}

// calculateSquare вычисляет квадрат числа и передаёт его в канал
func calculateSquare(num int, squareChan chan int) {
	square := num * num
	squareChan <- square
}

// printFromChanSequentially выводит результаты из канала в том же процессе, который её вызвал
func printFromChanSequentially(nPrints int, squareChan chan int) {
	for i := 0; i < nPrints; i++ {
		square := <-squareChan
		fmt.Println(square)
	}
}
