package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temperatureStream := make(chan float64)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	groupedTemperatures := make(map[int][]float64)

	go sendTemperatures(temperatures, temperatureStream)

	workers := 3
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go processTemperatures(temperatureStream, &wg, &mutex, groupedTemperatures)
	}

	wg.Wait()

	fmt.Println("Группировка температур:")
	for group, temps := range groupedTemperatures {
		fmt.Printf("%d: %v\n", group, temps)
	}
}

func sendTemperatures(temperatures []float64, out chan float64) {
	defer close(out)
	for _, temp := range temperatures {
		out <- temp
	}
}

func processTemperatures(in chan float64, wg *sync.WaitGroup, m *sync.Mutex, groupedTemperatures map[int][]float64) {
	defer wg.Done()
	for temp := range in {
		group := int(math.Round(temp/10.0)) * 10
		m.Lock()
		if _, ok := groupedTemperatures[group]; !ok {
			groupedTemperatures[group] = []float64{}
		}
		groupedTemperatures[group] = append(groupedTemperatures[group], temp)
		m.Unlock()
	}
}
