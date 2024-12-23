package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// producer generates random numbers and sends them to its output channel
func producer(id int, count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			num := rand.Intn(100) // Generate a random number
			fmt.Printf("Producer %d produced: %d\n", id, num)
			out <- num
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // Simulate work
		}
	}()
	return out
}

// fanIn combines multiple input channels into a single output channel
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Start a goroutine for each input channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	// Close the output channel once all input channels are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	// Create multiple producers
	producer1 := producer(1, 5)
	producer2 := producer(2, 5)
	producer3 := producer(3, 5)

	// Combine outputs from all producers using fanIn
	aggregated := fanIn(producer1, producer2, producer3)

	// Process aggregated results
	fmt.Println("Aggregated results:")
	for result := range aggregated {
		fmt.Println(result)
	}
}
