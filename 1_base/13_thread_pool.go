package main

import (
	"fmt"
	"time"
)

func threading1() {
	t := time.Now()

	const jobsCount = 15
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	go worker1(1, jobs, results)

	for i := 0; i < jobsCount; i += 1 {
		jobs <- i + 1
	}
	close(jobs)
	for i := 0; i < jobsCount; i += 1 {
		fmt.Printf("result %d : value = %d\n", i+1, <-results)
	}

	fmt.Printf("TIME ELAPSED: " + time.Since(t).String())
}

func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Worker %d finished\n", id)
		results <- j * j
	}
}

func threading2() {
	t := time.Now()

	const jobsCount, workerCount = 15, 3
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	for i := 0; i < workerCount; i += 1 {
		go worker2(i+1, jobs, results)
	}

	for i := 0; i < jobsCount; i += 1 {
		jobs <- i + 1
	}
	close(jobs)
	for i := 0; i < jobsCount; i += 1 {
		fmt.Printf("result %d : value = %d\n", i+1, <-results)
	}

	fmt.Printf("TIME ELAPSED: " + time.Since(t).String())
}

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Worker %d finished\n", id)
		results <- j * j
	}
}

func main() {
	//threading1()
	threading2()
}
