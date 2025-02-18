package main

import (
	"fmt"
	"sync"
)

// Producer generates data from a specified slice and sends it to the shared channel
func producer[T any](data []T, dataCh chan<- T, stopCh <-chan struct{}) {
	defer close(dataCh)
	for _, item := range data {
		select {
		case <-stopCh:
			fmt.Println("Producer stopping...")
			return
		default:
			dataCh <- item
			fmt.Println(fmt.Sprintf("Produced: %s", item))
		}
	}
}

// Consumer processes data from the shared channel
func consumer[T any](id int, dataCh <-chan T, wg *sync.WaitGroup, stopCh <-chan struct{}, errorCh chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println(fmt.Sprintf("Consumer %d stopping...", id))
			return
		case item, ok := <-dataCh:
			if !ok {
				fmt.Println(fmt.Sprintf("Consumer %d finished (channel closed).", id))
				return
			}

			// Simulate error handling
			s := any(item).(string)
			if s == "fake" {
				err := fmt.Errorf("consumer %d encountered an error with '%s'", id, item)
				errorCh <- err
				return
			}
			fmt.Println(fmt.Sprintf("Consumer %d processed: %s", id, item))
		}
	}
}

func main() {
	// Input data slice
	data := []any{"apple", "banana", "kiwi", "pear", "grape", "fake", "mango", "orange"}

	// Shared channels
	dataCh := make(chan any, len(data))
	stopCh := make(chan struct{})
	errorCh := make(chan error, 1) // Buffered channel to capture errors

	// WaitGroup for consumers
	var wg sync.WaitGroup

	// Start the producer
	go producer(data, dataCh, stopCh)

	// Start multiple consumers
	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, dataCh, &wg, stopCh, errorCh)
	}

	// Monitor for errors or completion
	go func() {
		// Wait for all consumers to complete
		wg.Wait()
		close(errorCh)
	}()

	// Stop all processing on error
	select {
	case err := <-errorCh:
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			close(stopCh) // Signal everyone to stop
		}
	case <-stopCh:
		fmt.Println("Stop signal received. Stopping.")
	}

	// Wait for consumers to clean up
	wg.Wait()
	fmt.Println("All consumers have stopped. Program exiting.")
}
