package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing", job)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)
	go Worker(1, jobs, &wg)
	go Worker(2, jobs, &wg)
	go Worker(3, jobs, &wg)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
}
