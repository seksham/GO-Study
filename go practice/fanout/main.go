package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a task to be processed
type Job struct {
	ID    int
	Image string // Simulated image, just a string for simplicity
}

// Result represents the outcome of processing a job
type Result struct {
	JobID          int
	ProcessedImage string
}

// worker processes jobs by counting words in the text
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d starting to process image %d\n", id, job.ID)

		// Fixed processing time of 2 seconds
		time.Sleep(2 * time.Second)

		processedImage := fmt.Sprintf("enhanced_%d.jpg", job.ID)
		fmt.Printf("Worker %d finished processing image %d\n", id, job.ID)
		results <- Result{JobID: job.ID, ProcessedImage: processedImage}
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Create jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Image: fmt.Sprintf("image_%d.jpg", j)}
	}
	close(jobs)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	startTime := time.Now()
	for result := range results {
		fmt.Println("Time: ", time.Since(startTime))
		fmt.Printf("Result: Image %d processed as %s\n", result.JobID, result.ProcessedImage)
	}

	fmt.Printf("All images processed in %v\n", time.Since(startTime))
}
