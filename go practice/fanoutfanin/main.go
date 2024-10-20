package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Image string
}

type Result struct {
	JobID          int
	ProcessedImage string
}

func worker(id int, jobs <-chan Job) <-chan Result {
	results := make(chan Result)
	go func() {
		for job := range jobs {
			fmt.Printf("Worker %d starting to process image %d\n", id, job.ID)
			time.Sleep(2 * time.Second)
			processedImage := fmt.Sprintf("enhanced_%d.jpg", job.ID)
			fmt.Printf("Worker %d finished processing image %d\n", id, job.ID)
			results <- Result{JobID: job.ID, ProcessedImage: processedImage}
		}
		close(results)
	}()
	return results
}

func fanIn(channels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	multiplexed := make(chan Result)

	multiplex := func(c <-chan Result) {
		defer wg.Done()
		for result := range c {
			multiplexed <- result
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)

	// Create jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Image: fmt.Sprintf("image_%d.jpg", j)}
	}
	close(jobs)

	// Start workers and collect their result channels
	var resultChannels []<-chan Result
	for i := 0; i < numWorkers; i++ {
		resultChannels = append(resultChannels, worker(i, jobs))
	}

	// Fan-in the result channels
	results := fanIn(resultChannels...)

	// Process results
	startTime := time.Now()
	for result := range results {
		fmt.Println("Time: ", time.Since(startTime))
		fmt.Printf("Result: Image %d processed as %s\n", result.JobID, result.ProcessedImage)
	}

	fmt.Printf("All images processed in %v\n", time.Since(startTime))
}