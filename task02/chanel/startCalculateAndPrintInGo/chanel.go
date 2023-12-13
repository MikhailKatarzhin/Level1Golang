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

	// Запуск из главного процесса отдельного процесса, запускающего вычисления квадратов чисел в горутинах
	startCalculateSquareInGoroutine(numbers, squareChan)

	// Запуск главным процессом процессов, выводящих результаты из канала, в горутинах
	printFromChanInGoroutine(len(numbers), squareChan)
}

// startCalculateSquare запускает процессы вычисления квадратов чисел в горутинах
func startCalculateSquare(numbers []int, squareChan chan int) {
	for _, num := range numbers {
		go calculateSquare(num, squareChan)
	}
}

// startCalculateSquareInGoroutine запускает горутин, запускающий процессы вычисления квадратов чисел в горутинах
func startCalculateSquareInGoroutine(numbers []int, squareChan chan int) {
	go startCalculateSquare(numbers, squareChan)
}

// calculateSquare вычисляет квадрат числа и передаёт его в канал
func calculateSquare(num int, squareChan chan int) {
	square := num * num
	squareChan <- square
}

// printFromChanInGoroutine выводит результаты из канала в отдельных горутинах.
// Процесс, который вызвал функцию, ждёт завершения всех горутин
func printFromChanInGoroutine(nPrints int, squareChan chan int) {
	done := make(chan bool)
	defer close(done)

	for i := 0; i < nPrints; i++ {
		go func() {
			square := <-squareChan
			fmt.Println(square)
			done <- true
		}()
	}

	for i := 0; i < nPrints; i++ {
		<-done
	}
}
