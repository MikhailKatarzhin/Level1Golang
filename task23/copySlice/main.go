package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)

	fmt.Println(removeAtIndex(arr, 1))
}

// copy копирует элементы из одного слайса в другой с учётом их позиций
func removeAtIndex(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
