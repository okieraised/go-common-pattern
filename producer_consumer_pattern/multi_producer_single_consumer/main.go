package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	inputs := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputCh := make(chan string)
	resultCh := make(chan string)
	errorCh := make(chan error, 1)
	resultDone := make(chan struct{})

	numWorkers := 3
	var wg sync.WaitGroup

	// Start consumer (now with its own error handling)
	go consumer(ctx, resultCh, errorCh, resultDone)

	// Error handler (from any producer or consumer)
	go func() {
		err := <-errorCh
		cancel()
		fmt.Println("[ERROR]", err)
	}()

	// Start worker pool
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, i, inputCh, resultCh, errorCh, &wg)
	}

	// Feed input strings
	go func() {
		for _, str := range inputs {
			select {
			case <-ctx.Done():
				return
			case inputCh <- str:
			}
		}
		close(inputCh)
	}()

	// Wait for workers and close result channel
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Wait for consumer to finish
	<-resultDone
}

func worker(ctx context.Context, id int, inputCh <-chan string, resultCh chan<- string, errorCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case str, ok := <-inputCh:
			if !ok {
				return
			}

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
			if rand.Intn(10) == 0 {
				errorCh <- fmt.Errorf("worker-%d: failed to process '%s'", id, str)
				return
			}

			result := fmt.Sprintf("worker-%d processed: %s", id, str)
			resultCh <- result
		}
	}
}

func consumer(ctx context.Context, resultCh <-chan string, errorCh chan<- error, done chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			done <- struct{}{}
			return
		case result, ok := <-resultCh:
			if !ok {
				done <- struct{}{}
				return
			}

			fmt.Println("[RESULT]", result)

			// Simulate consumer failure
			if rand.Intn(15) == 0 {
				errorCh <- fmt.Errorf("consumer: failed to handle '%s'", result)
				done <- struct{}{}
				return
			}
		}
	}
}
