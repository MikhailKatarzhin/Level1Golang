package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go worker(ch)

	time.Sleep(5 * time.Second)

	// Отправка в канал структуры, как сигнала для остановки
	ch <- struct{}{}
	close(ch)

	fmt.Println("Основной процесс отменил контекст, отправленный в горутину")
}

func worker(stopChan <-chan struct{}) {
	fmt.Println("Горутина начала работу")

	for {
		select {

		case <-stopChan: // При получении сигнала остановки в виде структуры из канала  уведомляет в консоль и прекращает работу
			fmt.Println("\nГорутина получила сигнал об остановке")
			return

		default: // Имитация работы
			fmt.Printf("\n[%s] Горутина работает", time.Now())
		}
	}
}
