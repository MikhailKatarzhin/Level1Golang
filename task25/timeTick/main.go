package main

import (
	"fmt"
	"time"
)

func main() {
	const nSec = 3

	fmt.Printf("Program sleep  %d seconds\n", nSec)

	start := time.Now()
	sleepSeconds(nSec)

	fmt.Printf("Program finished after: %d nanoseconds", time.Since(start))
}

func sleepSeconds(seconds int) {
	<-time.Tick(time.Duration(seconds) * time.Second)
}
