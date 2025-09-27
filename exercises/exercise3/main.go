package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
)

type Task struct {
	Name string
}

func Worker(id int, chanTask <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range chanTask {
		taskColor := Reset
		if task.Name == "Emergency" {
			taskColor = Red
		}
		fmt.Printf("Worker %s%d%s: Processing %sTask%s: %s%s%s\n",
			Red, id, Reset, Yellow, Reset, taskColor, task.Name, Reset)
		time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
	}
}

func Simulation(workers int, jobList []Task) {
	chanTask := make(chan Task)
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go Worker(i, chanTask, &wg)
	}

	for _, job := range jobList {
		chanTask <- job
	}

	close(chanTask)

	wg.Wait()
	fmt.Printf("%sAll tasks have completed%s\n\n", Yellow, Reset)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	job := []Task{
		{"Delivery"},
		{"Emergency"},
		{"Pickup"},
	}

	Simulation(2, job)

}
