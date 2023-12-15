package main

import (
	"fmt"
	"time"
)

func main() {
	nElement := 50

	start := time.Now()

	printFromChan(
		doubledIntFromChan(
			arrayToChan(
				makeFulledArray(nElement),
			),
		),
	)

	fmt.Printf("\nВыполнено за %d нс", time.Since(start).Nanoseconds())
}

// Запускает горутин, который Отправляет из массива числа в канал
// возвращает канал для получения чисел
func arrayToChan(numbers []int) <-chan int {
	numberChan := make(chan int)

	go func() {
		for _, num := range numbers {
			numberChan <- num
		}

		close(numberChan)
	}()

	return numberChan
}

// square Запускает горутин, который
// берёт число из канала с числами из массива,
// умножает их на 2
// отправляет в канал для удвоенных чисел
// возвращает канал для получения удвоенных чисел
func doubledIntFromChan(numberChan <-chan int) <-chan int {
	doubledChan := make(chan int)

	go func() {
		for num := range numberChan {
			doubled := num * 2
			doubledChan <- doubled
		}

		close(doubledChan)
	}()

	return doubledChan
}

// printFromChan Читает данные из канала и печатает их в std
func printFromChan(doubledChan <-chan int) {
	for square := range doubledChan {
		fmt.Println(square)
	}
}

// Создает массив чисел с указанным количеством элементов
// и заполняет его значением индекса элемента
func makeFulledArray(count int) []int {
	arr := make([]int, count)

	for i := 0; i < count; i++ {
		arr[i] = i
	}

	return arr
}
