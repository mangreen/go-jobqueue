package main

import (
	"fmt"
	"time"
)

func worker(jobChan <-chan int) {
	for j := range jobChan {
		fmt.Println("current job:", j)
		time.Sleep(3 * time.Second)
		fmt.Println("finished job:", j)
	}
}

func enqueue(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}

func main() {
	// make a channel with a capacity of 1
	jobChan := make(chan int, 1)

	// start the worker
	go worker(jobChan)

	// enqueue a job
	jobChan <- 1
	fmt.Println("enqueueed the job 1")
	jobChan <- 2
	fmt.Println("enqueueed the job 2")
	go func() {
		jobChan <- 3
		fmt.Println("enqueueed the job 3")
	}()

	fmt.Println("waiting the jobs")
	time.Sleep(10 * time.Second)
}
