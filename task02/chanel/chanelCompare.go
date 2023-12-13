package main

import (
	"fmt"
	"time"
)

func main() {
	//numbers := []int{2, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10, 4, 6, 8, 10}
	numbers := []int{2, 4, 6, 8, 10}

	start := time.Now()
	calculateAndPrintArraySquares(numbers, startCalculateSquare, printFromChanSequentially)
	fmt.Println("Время выполнения: ", time.Since(start).Microseconds())

	start = time.Now()
	calculateAndPrintArraySquares(numbers, startCalculateSquareInGoroutine, printFromChanSequentially)
	fmt.Println("Время выполнения: ", time.Since(start).Microseconds())

	start = time.Now()
	calculateAndPrintArraySquares(numbers, startCalculateSquareInGoroutine, printFromChanInGoroutine)
	fmt.Println("Время выполнения: ", time.Since(start).Microseconds())
}

func calculateAndPrintArraySquares(
	numbers []int,
	startCalculate func([]int, chan int),
	printFromChan func(int, chan int),
) {
	squareChan := make(chan int)
	defer close(squareChan)

	startCalculate(numbers, squareChan)

	printFromChan(len(numbers), squareChan)
}

func startCalculateSquare(numbers []int, squareChan chan int) {
	for _, num := range numbers {
		go calculateSquare(num, squareChan)
	}
}

func startCalculateSquareInGoroutine(numbers []int, squareChan chan int) {
	go startCalculateSquare(numbers, squareChan)
}

func calculateSquare(num int, squareChan chan int) {
	square := num * num
	squareChan <- square
}

func printFromChanSequentially(nPrints int, squareChan chan int) {
	for i := 0; i < nPrints; i++ {
		square := <-squareChan
		fmt.Println(square)
	}
}

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
