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

func Worker(id int, chanTask <-chan Task, done <-chan struct{}, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case task, ok := <-chanTask:
			if !ok {
				return
			}
			taskColor := Reset
			if task.Name == "Emergency" {
				taskColor = Red
			}
			fmt.Printf("Worker %s%d%s: Processing %sTask%s: %s%s%s\n",
				Red, id, Reset, Yellow, Reset, taskColor, task.Name, Reset)

			if task.Name == "Emergency" && rand.Intn(2) == 0 {
				errChan <- fmt.Errorf("Task %s failed", task.Name)
				continue
			}

			time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
		}
	}
}

func Simulation(workers int, jobList []Task) {
	chanTask := make(chan Task)
	errChan := make(chan error, len(jobList))
	done := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go Worker(i, chanTask, done, errChan, &wg)
	}

	go func() {
		err := <-errChan
		fmt.Println("ERROR:", err)
		close(done)
	}()

	rand.Shuffle(len(jobList), func(i, j int) {
		jobList[i], jobList[j] = jobList[j], jobList[i]
	})
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

	for i := 1; i <= 10; i++ {
		Simulation(10, job)
	}

}
