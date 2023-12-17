package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const nVal = 200
	const target = 5
	s := quicksort(fillByInt(make([]int, nVal), nVal))
	fmt.Println(s)
	idx := biSearch(s, target)
	if idx == -1 {
		fmt.Printf("\nTarget[%d] not found", target)
	} else {
		fmt.Printf("\nTarget[%d] was found under index %d", target, idx)
	}
}

// biSearch ищет в массиве/слайсе элемент и возвращает его индекс
// в случае отсутствия цели возвращает -1
func biSearch(arr []int, target int) int {
	length := len(arr)
	if arr[0] == target {
		return 0
	}

	if arr[length-1] == target {
		return length - 1
	}

	mid := length / 2

	if arr[mid] == target {
		return mid
	}

	idx := -1
	if target < mid && target > arr[0] {
		idx = biSearch(arr[:mid], target)
	} else if target > mid && target < arr[length-1] {
		idx = biSearch(arr[mid+1:], target)
	}

	// Если элемент не найден, возвращаем -1
	return idx
}

func fillByInt(a []int, nValues int) []int {
	for i := 0; i < nValues; i++ {
		// Создание случайного значения от 0 до 50 и добавление его во множество
		a[i] = rand.Intn(50)
	}

	return a
}

// quicksort - sorts array with O(n*log(n))
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, pivot := 0, len(a)-1

	for left < pivot {
		if a[left] > a[pivot] {
			if left+1 == pivot {
				a[pivot], a[left] = a[left], a[pivot]
			} else {
				a[pivot], a[pivot-1], a[left] = a[left], a[pivot], a[pivot-1]
			}
			pivot--
		} else {
			left++
		}
	}

	quicksort(a[:pivot])
	quicksort(a[pivot+1:])

	return a
}
