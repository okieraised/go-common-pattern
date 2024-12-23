package main

import (
	"fmt"
	"sync"
)

func producer[T any](link chan<- T, inputs []T, id int, wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done()
	for _, item := range inputs {
		select {
		case <-stopCh:
			fmt.Println(fmt.Sprintf("Stop signal received. Producer #%d stopped", id))
			return
		default:
			link <- item
			fmt.Println(fmt.Sprintf("Produced: %s", item))
		}
	}
}

func consumer[T any](link <-chan T, id int, wg *sync.WaitGroup, stopCh <-chan struct{}, errorCh chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println(fmt.Sprintf("Stop signal received. Consumer #%d stopped", id))
			return
		case item, ok := <-link:
			if !ok {
				return
			}

			// simulate error
			if any(item).(string) == "" {
				err := fmt.Errorf("consumer %d encountered an error", id)
				errorCh <- err
				return
			}
			fmt.Println(fmt.Sprintf("Consumer %d: consumed item %s", id, item))
		}
	}
}

func main() {
	//var once sync.Once

	dataCh := make(chan string)
	stopCh := make(chan struct{})
	errorCh := make(chan error, 1)

	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	nProducer := 4
	nConsumer := 2

	// start producer
	for i := 0; i < nProducer; i++ {
		wp.Add(1)
		inputs := []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "",
			"a", "b", "c", "d", "e", "f", "g", "h", "",
			"a", "b", "c", "d", "e", "f", "g", "h", "",
		}
		go producer(dataCh, inputs, i, wp, stopCh)
	}

	go func() {
		defer close(dataCh)
		wp.Wait()
	}()

	// start consumer
	for i := 0; i < nConsumer; i++ {
		wc.Add(1)
		go consumer(dataCh, i, wc, stopCh, errorCh)
	}

	go func() {
		// Wait for all consumers to complete
		wc.Wait()
		wp.Wait()
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

	wc.Wait()
	fmt.Println("All consumers have stopped. Program exiting.")
}
