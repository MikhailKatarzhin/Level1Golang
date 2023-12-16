package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создание синхронизированной мапы для конкурентного доступа к не1
	var data sync.Map

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

				// Выполняем запись в мапу
				data.Store(key, n)
				value, _ := data.Load(key)

				fmt.Printf("Записано: [%s : %d]\n", key, value)
			}
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод содержимого
	printMap(&data)
}

func printMap(data *sync.Map) {
	fmt.Println("\nСодержимое map:")

	data.Range(printKeyVal)
}

func printKeyVal(key any, value any) bool {
	fmt.Printf(
		"[%s : %d]\n",
		key,
		value,
	)
	return true
}
