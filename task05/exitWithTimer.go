package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	nSeconds := 5 // Установите N на количество секунд до завершения программы

	readFromChannel(
		writeToChannelDuringTimer(
			getTimerChanForNSecond(
				nSeconds,
			),
		),
	)

	fmt.Println("Program finished.")
}

// getTimerChanForNSecond создает канал-таймер на nSeconds секунд
func getTimerChanForNSecond(nSeconds int) *time.Timer {
	return time.NewTimer(time.Duration(nSeconds) * time.Second)
}

// writeToChannelDuringTimer пишет данные в канал intChan пока не придёт что-либо из канала-таймера timerChan
func writeToChannelDuringTimer(timer *time.Timer) <-chan int {
	intChan := make(chan int)

	go func() {
		for i := 0; ; i++ {
			select {
			case <-timer.C:
				close(intChan)
				return
			default:
				if i == math.MaxInt {
					i = 0
				}
				intChan <- i
			}
		}
	}()

	return intChan
}

// readFromChannel читает данные из канала и выводит их в консоль, пока они поступают
func readFromChannel(ch <-chan int) {
	for v := range ch {
		fmt.Println("Received:", v)
	}
}
