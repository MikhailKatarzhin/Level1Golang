package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go worker(ctx, &wg)

	wg.Wait()
	fmt.Println("Основной процесс получила сигнал от горутины о завершении")
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("Горутина начала работу")
	for {
		select {
		case <-ctx.Done(): // Получение сигнала об истечении времени контекста
			fmt.Println("\nГорутина получила сигнал о завершении времени ожидания контекста")
			wg.Done()
			return
		default: // Имитация работы
			fmt.Printf("\n[%s] Горутина работает", time.Now())
		}
	}
}
