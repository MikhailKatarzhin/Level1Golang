package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	const countGo = 2000
	fmt.Printf("Число горутин: %d\n", countGo)

	// Запускаем несколько горутин для инкрементирования счётчика конкурентно
	for i := 0; i < countGo; i++ {
		wg.Add(1)

		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait() // Ждём завершения всех горутин

	fmt.Printf("Итоговое значение счётчика: %d", counter.value)
}
