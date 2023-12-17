package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)

	fmt.Println(removeAtIndex(arr, 1))
}

// заменяет удаляемый элемент последним и сокращает длину слайса
// Быстрее, но теряется упорядоченность
func removeAtIndex(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
