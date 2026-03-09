package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func worker(id int, ch chan int) {
	defer wg.Done()

	mutex.Lock()
	counter++
	fmt.Printf("Worker %d incremented counter to %d\n", id, counter)
	mutex.Unlock()

	ch <- id
}

func main() {
	ch := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch)
	}

	wg.Wait()
	close(ch)

	for id := range ch {
		fmt.Printf("Received from worker %d\n", id)
	}
	fmt.Println("Counter: ", counter)
}
