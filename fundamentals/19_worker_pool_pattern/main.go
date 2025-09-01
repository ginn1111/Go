package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID   int
	Edge int
}

type Result struct {
	JobId  int
	Square int
}

func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		results <- Result{job.ID, job.Edge * job.Edge}
	}

}

func collectResult(results chan Result, Done chan bool) {
	for res := range results {
		fmt.Println(res.JobId, res.Square)
	}

	Done <- true
}

func addJob(jobs chan<- Job, edge, id int) {
	jobs <- Job{ID: id, Edge: edge}
}

func main() {
	var WorkerCount = 3
	_ = WorkerCount
	var wg sync.WaitGroup

	jobTasks := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	results := make(chan Result, len(jobTasks))
	jobs := make(chan Job, len(jobTasks))

	go func() {
		for id, jobTask := range jobTasks {
			addJob(jobs, jobTask, id)
		}

		close(jobs)
	}()

	for _ = range WorkerCount {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	Done := make(chan bool)

	go collectResult(results, Done)

	wg.Wait()

	close(results)

	<-Done

}
