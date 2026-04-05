// Observer pattern is unnecessary in Go. Channels and goroutines naturally decouple
// producers from consumers, replacing the Subject/Observer mechanism idiomatically.
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// digitObserver reads numbers from a channel and prints them as digits.
func digitObserver(name string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Printf("%s:%d\n", name, n)
		time.Sleep(100 * time.Millisecond)
	}
}

// graphObserver reads numbers from a channel and prints them as a bar of '*'.
func graphObserver(name string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Printf("%s:%s\n", name, strings.Repeat("*", n))
		time.Sleep(100 * time.Millisecond)
	}
}

// generateRandomNumbers produces 20 random numbers (0-49) and fans them out
// to all subscriber channels.
func generateRandomNumbers(subscribers []chan<- int) {
	for i := 0; i < 20; i++ {
		n := rand.Intn(50)
		for _, ch := range subscribers {
			ch <- n
		}
	}
	for _, ch := range subscribers {
		close(ch)
	}
}

func main() {
	// Create channels for each observer (fan-out).
	digitCh := make(chan int)
	graphCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Launch observer goroutines.
	go digitObserver("DigitObserver", digitCh, &wg)
	go graphObserver("GraphObserver", graphCh, &wg)

	// Generate numbers and broadcast to all observers.
	generateRandomNumbers([]chan<- int{digitCh, graphCh})

	// Wait for observers to finish processing.
	wg.Wait()
}
