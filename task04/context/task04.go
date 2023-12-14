package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
)

func main() {
	numWorkerReaders, err := askInputNumWorkers()
	if err != nil {
		println(err)
		return
	}

	// Создаем контекст для завершения
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	// Передача сигнала по нажатию Ctrl+C
	signal.Notify(sigChan, os.Interrupt)

	//Горутин для отслеживания сигнала и завершения выполнения
	go func() {
		<-sigChan
		fmt.Println("\nReceived SIGINT. Shutting down...")
		cancel() // При получении сигнала отменяем контекст для завершения работы генератора чисел
		close(sigChan)
	}()

	startWorkerPull(numWorkerReaders, genIntsToChan(ctx))
}

func askInputNumWorkers() (int, error) {
	var numWorkerReaders int

	fmt.Println("Enter the number of worker readers:")
	_, err := fmt.Scan(&numWorkerReaders)

	if err != nil {
		return 0, fmt.Errorf(
			"can't scan input: Invalid input: %w", err,
		)
	}

	if numWorkerReaders <= 0 {
		fmt.Println("Number of worker readers should be a positive integer")
		return 0, fmt.Errorf(
			"invalid input: Number of worker readers should be a positive integer: %w", err,
		)
	}

	return numWorkerReaders, nil
}

func genIntsToChan(ctx context.Context) <-chan int {
	dataChan := make(chan int)

	// Горутина, генерирующая целые числа и отправляющая их в канал
	go func() {
		defer close(dataChan)

		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			case dataChan <- i:
				if i == math.MaxInt {
					i = 1
				}
			}
		}

	}()

	return dataChan
}

func startWorkerPull(nWorkers int, dataChan <-chan int) {
	var wg sync.WaitGroup

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go workerReader(i, dataChan, &wg)
	}

	wg.Wait() // Ожидание завершения всех воркеров
	fmt.Println("All workerReaders have finished.")
}

func workerReader(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataChan {
		fmt.Printf("workerReader [id:%d]: %d\n", id, data)
	}
}

/* После нажатия ctrl + c

Received SIGINT. Shutting down...
workerReader [id:0]: 15510
All workerReaders have finished.

*/
