package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(5 * time.Second)

	cancel()

	fmt.Println("Основной процесс отменил контекст, отправленный в горутину")
}

func worker(ctx context.Context) {
	fmt.Println("Горутина начала работу")
	for {
		select {
		case <-ctx.Done(): // При получении сигнала отмены контекста
			fmt.Println("\nГорутина получила сигнал об отмене контекста")
			return
		default: // Имитация работы
			fmt.Printf("\n[%s] Горутина работает", time.Now())
		}
	}
}
