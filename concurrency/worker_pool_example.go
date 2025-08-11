package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numOfJobs = 10
	const numOfWorkers = 3
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	var wg sync.WaitGroup

	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results, i)
	}

	//send jobs
	for i := 1; i <= numOfJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)
	fmt.Println("All jobs done. Displaying results.")
	for result := range results {
		fmt.Println("Result:", result)
	}
}

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- int, id int) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println(fmt.Sprintf("worker %d is processing job %d", id, job))
		time.Sleep(time.Millisecond * 500)
		results <- job
	}
}
