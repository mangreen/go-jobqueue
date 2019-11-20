package main

import (
	"fmt"
	"time"
)

func worker(jobChan <-chan int) {
	for j := range jobChan {
		fmt.Println("current job:", j)
	}
}

func main() {
	// make a channel with a capacity of 1024
	jobChan := make(chan int, 1024)

	// start the worker
	go worker(jobChan)

	// enqueue a job
	jobChan <- 1
	fmt.Println("enqueueed the job 1")
	jobChan <- 2
	fmt.Println("enqueueed the job 2")
	jobChan <- 3
	fmt.Println("enqueueed the job 3")

	time.Sleep(1 * time.Second)
}
