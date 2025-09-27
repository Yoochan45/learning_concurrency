package main

import (
	"fmt"
	"sync"
)

func PrintNumber(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(name, ":")
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go PrintNumber("goroutine A", &wg)
	go PrintNumber("goroutine B", &wg)

	wg.Wait()
	fmt.Println("All Done")
}
