package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const nVal = 200
	s := fillByInt(make([]int, nVal), nVal)
	fmt.Println(s)
	fmt.Println()
	fmt.Println(quicksort(s))
}

func fillByInt(a []int, nValues int) []int {
	for i := 0; i < nValues; i++ {
		// Создание случайного значения от 0 до 500 и добавление его во множество
		a[i] = rand.Intn(500)
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
