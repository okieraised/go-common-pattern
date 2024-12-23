package main

import (
	"fmt"
	"sync"
)

func producer[T any](link chan<- T, inputs []T, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range inputs {
		link <- msg
	}
	fmt.Println(fmt.Sprintf("Producer #%d finished", id))
}

func consumer[T any](link <-chan T, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range link {
		fmt.Println(fmt.Sprintf("Consumer #%d: received %s", id, msg))
	}
}

func main() {
	ch := make(chan string)
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	nProducer := 4
	nConsumer := 2

	for i := 0; i < nProducer; i++ {
		wp.Add(1)
		inputs := []string{
			"a", "b", "c", "d", "e", "f", "g", "h",
			"a", "b", "c", "d", "e", "f", "g", "h",
			"a", "b", "c", "d", "e", "f", "g", "h",
		}
		go producer(ch, inputs, i, wp)
	}

	for i := 0; i < nConsumer; i++ {
		wc.Add(1)
		go consumer(ch, i, wc)
	}

	wp.Wait()
	close(ch)
	wc.Wait()

	fmt.Println("completed")
}
