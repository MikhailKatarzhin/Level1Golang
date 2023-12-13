package main

import (
	"fmt"
	"time"
)

func main() {
	nElement := 500

	start := time.Now()
	fmt.Printf(
		"Сумма квадратов чисел: %d ",
		<-sum(
			square(
				arrayToChan(
					makeFulledArray(nElement),
				),
			),
		),
	)
	fmt.Printf("\nПосчитано за %d нс", time.Since(start).Nanoseconds())
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
// возводит их в квадрат
// и отправляет в канал для квадратов чисел
// возвращает канал для получения квадратов чисел
func square(numberChan <-chan int) <-chan int {
	squareChan := make(chan int)

	go func() {
		for num := range numberChan {
			square := num * num
			squareChan <- square
		}

		close(squareChan)
	}()

	return squareChan
}

// sum Запускает горутин, который
// берёт число из канала для квадратов чисел,
// добавляет их в общую сумму
// после получения всех квадратов отправляет сумму в канал для суммы
// возвращает канал для получения суммы
func sum(squareChan <-chan int) <-chan int {
	sumChan := make(chan int)

	go func() {
		sum := 0

		for square := range squareChan {
			sum += square
		}

		sumChan <- sum
		close(sumChan)
	}()

	return sumChan
}

// Создает массив чисел с указанным количеством элементов
// и заполняет его удвоенным значением индекса элемента
func makeFulledArray(count int) []int {
	arr := make([]int, count)

	for i := 0; i < count; i++ {
		arr[i] = i * 2
	}

	return arr
}
