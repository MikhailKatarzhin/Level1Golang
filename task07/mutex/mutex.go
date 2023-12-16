package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создание мапа и мьютекса для конкурентного доступа к нему
	data := make(map[string]int)
	var mutex sync.RWMutex

	// Количество горутин для записи в мап
	numGoRoutines := 10
	numRecords := 10

	// Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		go func(n int) {
			defer wg.Done()

			for y := 0; y < numRecords; y++ {
				//Подготовка ключа
				key := fmt.Sprintf("Key%d-%d", n, y)

				// Захватываем мьютекс перед записью в мапу
				mutex.Lock()

				// Выполняем запись в мапу
				data[key] = n

				// Отпускаем мьютекс сразу после записи
				mutex.Unlock()

				fmt.Printf("Записано: [%s : %d]\n", key, data[key])
			}
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод содержимого
	printMap(data)
}

func printMap(data map[string]int) {
	fmt.Println("\nСодержимое map:")

	for key, value := range data {
		fmt.Printf("[%s : %d]\n", key, value)
	}
}
