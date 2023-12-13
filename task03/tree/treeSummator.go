package main

import (
	"fmt"
	"time"
)

func main() {
	nElement := 500

	start := time.Now()
	fmt.Printf(
		"Сумма квадратов чисел: %d",
		treeSquaresSum(
			makeFulledArray(nElement),
		),
	)
	fmt.Printf("\nПосчитано за %d нс", time.Since(start).Nanoseconds())
}

// Вычисление суммы квадратов, получаемых из одного канала, через бинарное дерево
func treeSquaresSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0] * nums[0]
	}

	return sumFromChans(startLeftAndRightSum(nums))
}

// Запускает вычисление сумм двух половин одного слайса
func startLeftAndRightSum(nums []int) (<-chan int, <-chan int) {
	mid := len(nums) / 2
	return startTreeSquaresSum(nums[:mid]), startTreeSquaresSum(nums[mid:])
}

// Запускает горутин на вычисление суммы квадратов для указанного слайса
func startTreeSquaresSum(nums []int) <-chan int {
	sumChan := make(chan int)

	go func() {
		sumChan <- treeSquaresSum(nums)
		close(sumChan)
	}()

	return sumChan
}

// Складывает результаты выполнения двух каналов
func sumFromChans(first <-chan int, second <-chan int) int {
	return <-first + <-second
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
