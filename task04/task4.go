package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	const numWorkerReaders = 5

	sigChan := make(chan os.Signal)
	// Завершение программы по сигналу Ctrl+C
	signal.Notify(sigChan, syscall.SIGTERM)

	startWorkerPull(numWorkerReaders, genIntsToChan(sigChan))
}

func genIntsToChan(sigChan <-chan os.Signal) <-chan int {
	dataChan := make(chan int)

	// Горутин, генерирующий целочисленные и отправляет в канал
	go func() {
		for i := 1; i > 0; i++ {
			select {
			case <-sigChan:
				close(dataChan)
				i = 0
			default:
				dataChan <- i
				if i > 1000000 {
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
