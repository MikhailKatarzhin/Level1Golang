package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)

	fmt.Println(removeAtIndex(arr, 1))
}

// append изменяет длину и ёмкость слайса
func removeAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
