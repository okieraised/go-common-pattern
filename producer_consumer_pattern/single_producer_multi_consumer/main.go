package main

import (
	"fmt"
	"sync"
)

func producer[T any](ch chan T, inputs []T) {
	defer close(ch)
	for _, param := range inputs {
		ch <- param
	}
}

func consumer[T any](idx int, ch <-chan T, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i, ok := <-ch
		if ok {
			fmt.Println(fmt.Sprintf("Consumer #%d: received %s", idx, i))
		} else {
			fmt.Println(fmt.Sprintf("Consumer #%d: no more values to process, exiting", idx))
			return
		}
	}
}

func main() {

	inputs := []string{
		"a", "b", "c", "d", "e", "f", "g", "h",
		"a", "b", "c", "d", "e", "f", "g", "h",
		"a", "b", "c", "d", "e", "f", "g", "h",
	}

	ch := make(chan string) // both block or non-block are ok
	var wg sync.WaitGroup
	consumerCount := 10

	go producer(ch, inputs)

	for i := 1; i <= consumerCount; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	fmt.Println("Main waiting")
	wg.Wait()
	fmt.Println("Main exiting")
}
